package queue

import (
	"container/list"
	"modules/world"
	"modules/world/bus"
	"sync"
	"time"
)

type Queue struct {
	Locker  sync.RWMutex
	Pending bool
	List    *list.List
	Timer   *time.Timer
}

func (c *Queue) InstallTo(b *bus.Bus) {
	b.BindCloseEvent(c, c.close)
	b.DoSendToQueue = b.WrapHandleSend(c.Append)
	b.DoDiscardQueue = b.Wrap(c.Flush)
}

func (c *Queue) close(b *bus.Bus) {
	c.Flush(b)
}
func (c *Queue) Flush(b *bus.Bus) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	timer := c.Timer
	if timer != nil {
		timer.Stop()
	}
	c.List.Init()
	c.Pending = false
}
func (c *Queue) Append(b *bus.Bus, cmd *world.Command) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	if c.List.Len() > 0 || c.Pending {
		c.List.PushBack(cmd)
		return
	}
	c.exec(b, cmd)
}
func (c *Queue) check(b *bus.Bus) {
	if c.List.Len() != 0 && !c.Pending {
		e := c.List.Front()
		c.List.Remove(e)
		cmd := e.Value.(*world.Command)
		c.exec(b, cmd)
	}

}
func (c *Queue) AfterDelay(b *bus.Bus) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Pending = false
	c.Timer = nil
	c.check(b)
}
func (c *Queue) exec(b *bus.Bus, cmd *world.Command) {
	b.DoSend(cmd)
	delay := b.GetQueueDelay()
	if delay >= 0 {
		c.Pending = true
		c.Timer = time.AfterFunc(time.Duration(delay)*time.Millisecond, func() { c.AfterDelay(b) })
	}
}

func New() *Queue {
	q := &Queue{}
	q.List = list.New()
	q.List.Init()
	return q
}
