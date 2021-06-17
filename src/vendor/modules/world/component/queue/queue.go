package queue

import (
	"container/list"
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
	b.BindReadyEvent(c, c.Ready)
	b.BindCloseEvent(c, c.close)
}
func (c *Queue) Ready(b *bus.Bus) {
	c.List = list.New()
	c.List.Init()
}
func (c *Queue) close(b *bus.Bus) {
	c.Flush()
}
func (c *Queue) Flush() {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	timer := c.Timer
	if timer != nil {
		timer.Stop()
	}
	c.List.Init()
	c.Pending = false
}
func (c *Queue) Append(b *bus.Bus, msg []byte) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	if c.List.Len() > 0 || c.Pending {
		c.List.PushBack(msg)
		return
	}
	c.exec(b, msg)
}
func (c *Queue) check(b *bus.Bus) {
	if c.List.Len() != 0 && !c.Pending {
		e := c.List.Front()
		c.List.Remove(e)
		msg := e.Value.([]byte)
		c.exec(b, msg)
	}

}
func (c *Queue) AfterDelay(b *bus.Bus) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Pending = false
	c.Timer = nil
	c.check(b)
}
func (c *Queue) exec(b *bus.Bus, msg []byte) {
	b.DoSend([]byte(msg))
	delay := b.GetQueueDelay()
	if delay >= 0 {
		c.Pending = true
		c.Timer = time.AfterFunc(time.Duration(delay)*time.Millisecond, func() { c.AfterDelay(b) })
	}
}

func New() *Queue {
	return &Queue{}
}
