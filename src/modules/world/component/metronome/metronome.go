package metronome

import (
	"container/list"
	"hellclient/modules/world"
	"hellclient/modules/world/bus"
	"sync"
	"time"
)

var DefaultCheckInterval = 50 * time.Millisecond
var DefaultTick = time.Second

var DefaultBeats = 10

type Metronome struct {
	Locker   sync.Mutex
	tick     time.Duration
	beats    int
	queue    *list.List
	sent     *list.List
	interval time.Duration
	tickerC  chan int
}

func (m *Metronome) Tick() time.Duration {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.tick
}

func (m *Metronome) SetTick(t time.Duration) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.tick = t
}
func (m *Metronome) Beats() int {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.beats
}

func (m *Metronome) SetBeats(b int) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.beats = b
}
func (m *Metronome) getBeats() int {
	if m.beats > 0 {
		return m.beats
	}
	return 1
}
func (m *Metronome) Reset(b *bus.Bus) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.sent = list.New()
	go m.play(b)
}
func (m *Metronome) Space() int {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.clean()
	space := m.getBeats() - m.sent.Len()
	if space < 0 {
		space = 0
	}
	return space
}
func (m *Metronome) Queue() []string {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	result := make([]string, m.queue.Len())
	for e := m.queue.Front(); e != nil; e = e.Next() {
		cmds := e.Value.([]*world.Command)
		for k := range cmds {
			result = append(result, cmds[k].Mesasge)
		}
	}
	return result
}
func (m *Metronome) Discard(force bool) bool {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.discard(force)
}
func (m *Metronome) discard(force bool) bool {
	var result = m.queue.Len()
	q := list.New()
	if !force {
		for e := m.queue.Front(); e != nil; e = e.Next() {
			cmds := e.Value.([]*world.Command)
			c := []*world.Command{}
			for _, v := range cmds {
				if v.Locked {
					c = append(c, v)
				}
			}
			if len(c) > 0 {
				q.PushBack(c)
			}
		}
	}
	m.queue = q
	return result != m.queue.Len()
}
func (m *Metronome) LockQueue() {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	for e := m.queue.Front(); e != nil; e = e.Next() {
		cmds := e.Value.([]*world.Command)
		for _, v := range cmds {
			v.Locked = true
		}
	}
}
func (m *Metronome) Full() {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.full()
}
func (m *Metronome) full() {
	t := time.Now()
	b := m.getBeats()
	m.sent = list.New()
	for i := 0; i < b; i++ {
		m.sent.PushBack(&t)
	}
}
func (m *Metronome) FullTick() {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.fullTick()
}
func (m *Metronome) fullTick() {
	t := time.Now()
	b := m.getBeats()
	for i := m.sent.Len(); i < b; i++ {
		m.sent.PushBack(&t)
	}
}

func (m *Metronome) Send(bus *bus.Bus, cmd *world.Command) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	t := time.Now()
	m.sent.PushBack(&t)
	bus.DoSend(cmd)
}
func (m *Metronome) append(rawcmds []*world.Command, grouped bool) {
	cmds := make([]*world.Command, 0, len(rawcmds))
	for k := range rawcmds {
		cmds = append(cmds, rawcmds[k].Split("\n")...)
	}
	if grouped {
		m.queue.PushBack(cmds)
	} else {
		for k := range cmds {
			m.queue.PushBack([]*world.Command{cmds[k]})
		}
	}
}
func (m *Metronome) clean() {
	t := time.Now()
	for e := m.sent.Front(); e != nil; e = e.Next() {
		sent := e.Value.(*time.Time)
		if t.Sub(*sent) > m.tick {
			m.sent.Remove(e)
		}
	}
}
func (m *Metronome) Play(bus *bus.Bus) {
	go m.play(bus)
}
func (m *Metronome) play(bus *bus.Bus) {
	if !bus.GetConnConnected() {
		return
	}
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.clean()
	b := m.getBeats()
	for m.queue.Len() != 0 && m.sent.Len() < b {
		e := m.queue.Front()
		cmds := e.Value.([]*world.Command)
		if b-m.sent.Len() < len(cmds) {
			//避免cmds长于beats时永远不发送
			if m.sent.Len() != 0 {
				return
			}
		}
		m.queue.Remove(e)
		for k := range cmds {
			t := time.Now()
			m.sent.PushBack(&t)
			bus.DoSend(cmds[k])
		}
	}
}

func (m *Metronome) Start(b *bus.Bus) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.startTicker(b)
}
func (m *Metronome) Stop(b *bus.Bus) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.stopTicker()
}
func (m *Metronome) SetInterval(b *bus.Bus, interval time.Duration) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.interval = interval
	m.startTicker(b)
}
func (m *Metronome) Interval() time.Duration {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.interval
}
func (m *Metronome) startTicker(b *bus.Bus) {
	m.stopTicker()
	go func() {
		interval := m.interval
		if interval <= 0 {
			interval = DefaultCheckInterval
		}
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case <-m.tickerC:
				return
			case <-t.C:
				go m.Play(b)
			}
		}
	}()
}
func (m *Metronome) stopTicker() {
	if m.tickerC != nil {
		close(m.tickerC)
	}
}
func (m *Metronome) InstallTo(b *bus.Bus) {
	m.SetInterval(b, DefaultCheckInterval)

	b.GetMetronomeBeats = m.Beats
	b.SetMetronomeBeats = m.SetBeats
	b.DoResetMetronome = b.Wrap(m.Reset)
	b.GetMetronomeSpace = m.Space
	b.GetMetronomeQueue = m.Queue
	b.DoDiscardMetronome = m.Discard
	b.DoLockMetronomeQueue = m.LockQueue
	b.DoFullMetronome = m.Full
	b.DoFullTickMetronome = m.FullTick
	b.SetMetronomeInterval = b.WrapSetDuration(m.SetInterval)
	b.GetMetronomeInterval = m.Interval
	b.SetMetronomeTick = m.SetTick
	b.GetMetronomeTick = m.Tick
	b.DoPushMetronome = b.WrapHandlePushGroupedCommands(m.Push)
	b.DoMetronomeSend = b.WrapHandleSend(m.Send)
	b.BindCloseEvent(m, m.Stop)
}
func (m *Metronome) Push(b *bus.Bus, cmds []*world.Command, grouped bool) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.append(cmds, grouped)
	go m.play(b)
}
func New() *Metronome {
	return &Metronome{
		tick:  DefaultTick,
		beats: DefaultBeats,
		queue: list.New(),
		sent:  list.New(),
	}
}
