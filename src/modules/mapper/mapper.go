package mapper

import (
	"sync"
)

type Path struct {
	Command     string
	Delay       int
	From        string
	To          string
	Tags        map[string]bool
	ExcludeTags map[string]bool
}

func NewPath() *Path {
	return &Path{
		Tags:        map[string]bool{},
		ExcludeTags: map[string]bool{},
	}
}

type Room struct {
	ID    string
	Name  string
	Exits []*Path
}

func NewRoom() *Room {
	return &Room{
		Exits: []*Path{},
	}
}

type Mapper struct {
	Locker sync.RWMutex
	rooms  *Rooms
	tags   map[string]bool
	fly    []*Path
}

func (m *Mapper) FlyList() []*Path {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return append(m.fly)
}
func (m *Mapper) SetFlyList(fly []*Path) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.fly = fly
}
func (m *Mapper) FlashTags() {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.tags = map[string]bool{}
}
func (m *Mapper) AddTags(tagnames []string) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	for _, v := range tagnames {
		m.tags[v] = true
	}
}
func (m *Mapper) SetTag(tagname string, enabled bool) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	if enabled {
		m.tags[tagname] = true
	} else {
		delete(m.tags, tagname)
	}
}
func (m *Mapper) Tags() []string {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	var result = make([]string, 0, len(m.tags))
	for k, v := range m.tags {
		if v {
			result = append(result, k)
		}
	}
	return result
}
func (m *Mapper) WalkAll(targets []string, fly bool, max_distance int) *WalkAllResult {
	a := NewWalkAll()
	a.rooms = m.rooms
	a.tags = m.tags
	a.fly = m.fly
	a.Targets = targets
	a.MaxDistance = max_distance
	if !fly {
		a.fly = nil
	}
	return a.Start()
}
func (m *Mapper) GetPath(from string, fly bool, to []string) []*Step {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	w := m.newWalking()
	w.from = from
	w.to = to
	if !fly {
		w.fly = nil
	}
	return w.Walk()
}
func (m *Mapper) newWalking() *Walking {
	walking := NewWalking()
	walking.rooms = m.rooms
	walking.tags = m.tags
	walking.fly = m.fly
	return walking
}

func (m *Mapper) GetRoomID(name string) []string {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.rooms.GetRoomID(name)
}

func (m *Mapper) GetRoomName(id string) string {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.rooms.GetRoomName(id)
}
func (m *Mapper) SetRoomName(id string, name string) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.rooms.SetRoomName(id, name)
}
func (m *Mapper) AddPath(id string, p *Path) bool {
	if p.Command == "" {
		return false
	}
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.rooms.AddPath(id, p)
}
func (m *Mapper) AddTemporaryPath(id string, p *Path) bool {
	if p.Command == "" {
		return false
	}
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.rooms.AddTemporaryPath(id, p)
}

func (m *Mapper) ClearRoom(id string) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.rooms.ClearRoom(id)
}

func (m *Mapper) NewArea(size int) []string {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.rooms.NewArea(size)
}
func (m *Mapper) GetExits(id string, all bool) []*Path {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.rooms.GetExits(id, m.tags, all)

}
func (m *Mapper) Reset() {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.tags = map[string]bool{}
	m.fly = []*Path{}
	m.rooms.Reset()
}
func (m *Mapper) ResetTemporary() {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.rooms.ResetTemporary()
}

func New() *Mapper {
	return &Mapper{
		rooms: NewRooms(),
		tags:  map[string]bool{},
		fly:   []*Path{},
	}
}
