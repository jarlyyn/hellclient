package queue

import (
	"container/list"
	"hellclient/modules/world"
	"hellclient/modules/world/bus"
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
	b.DoDiscardQueue = b.WrapDiscard(c.Flush)
	b.DoLockQueue = c.LockQueue
	b.GetQueue = c.ListQueue
	b.BindQueueDelayUpdatedEvent(c, c.onDelayUpdate)
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
	c.Flush(b, true)
}
func (c *Queue) Flush(b *bus.Bus, force bool) int {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	var l int
	if force {
		l = c.List.Len()
		c.List.Init()
	} else {
		var result = list.New()
		for i := c.List.Front(); i != nil; i = i.Next() {
			c := i.Value.(*world.Command)
			if c.Locked {
				result.PushBack(c)
			} else {
				l++
			}
		}
		c.List = result
	}
	timer := c.Timer
	if c.List.Len() == 0 {
		if timer != nil {
			timer.Stop()
		}
		c.Pending = false
	} else {
		c.Append(b, nil)
	}
	return l
}
func (c *Queue) LockQueue() {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	for i := c.List.Front(); i != nil; i = i.Next() {
		c := i.Value.(*world.Command)
		c.Locked = true
	}
}

func (c *Queue) Append(b *bus.Bus, cmd *world.Command) {
	c.Locker.Lock()
	cmds := cmd.Split("\n")
	for _, v := range cmds {
		c.List.PushBack(v)
	}
	if !c.Pending {
		c.delay(b)
	}
	c.Locker.Unlock()
}

func (c *Queue) AfterDelay(b *bus.Bus) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.send(b)
	if c.List.Len() == 0 {
		c.Pending = false
		if c.Timer != nil {
			c.Timer.Stop()
		}
		c.Timer = nil
	}
}
func (c *Queue) send(b *bus.Bus) {
	if c.List.Len() != 0 {
		e := c.List.Front()
		c.List.Remove(e)
		cmd := e.Value.(*world.Command)
		b.DoMetronomeSend(cmd)
		if c.List.Len() != 0 {
			c.delay(b)
		}
	}

}
func (c *Queue) delay(b *bus.Bus) {
	delay := b.GetQueueDelay()
	if delay > 0 {
		c.Pending = true
		c.Timer = time.AfterFunc(time.Duration(delay)*time.Millisecond, func() {
			c.AfterDelay(b)
		})
	} else {
		c.send(b)
	}
}

func (c *Queue) onDelayUpdate(b *bus.Bus) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Pending = false
	if c.Timer != nil {
		c.Timer.Stop()
		c.Timer = nil
	}
	c.delay(b)

}
func New() *Queue {
	q := &Queue{}
	q.List = list.New()
	q.List.Init()
	return q
}
