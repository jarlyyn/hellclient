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
	b.DoDiscardQueue = c.Flush
	b.GetQueue = c.ListQueue
}
func (c *Queue) ListQueue() []*world.Command {
	c.Locker.RLock()
	defer c.Locker.RUnlock()
	var result = make([]*world.Command, 0, c.List.Len())
	for i := c.List.Front(); i != nil; i = i.Next() {
		c := i.Value.(*world.Command)
		result = append(result, c)
	}
	return result
}
func (c *Queue) close(b *bus.Bus) {
	c.Flush()
}
func (c *Queue) Flush() int {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	timer := c.Timer
	if timer != nil {
		timer.Stop()
	}
	l := c.List.Len()
	c.List.Init()
	c.Pending = false
	return l
}
func (c *Queue) Append(b *bus.Bus, cmd *world.Command) {
	c.Locker.Lock()
	if c.List.Len() > 0 || c.Pending {
		c.List.PushBack(cmd)
		c.Locker.Unlock()
		return
	}
	c.Locker.Unlock()
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
	c.Pending = false
	c.Timer = nil
	c.Locker.Unlock()
	c.check(b)
}
func (c *Queue) exec(b *bus.Bus, cmd *world.Command) {
	delay := b.GetQueueDelay()
	if delay < 1 {
		delay = 1
	}
	c.Locker.Lock()
	c.Pending = true
	c.Timer = time.AfterFunc(time.Duration(delay)*time.Millisecond, func() { c.AfterDelay(b) })
	c.Locker.Unlock()
	b.DoSend(cmd)
}

func New() *Queue {
	q := &Queue{}
	q.List = list.New()
	q.List.Init()
	return q
}
