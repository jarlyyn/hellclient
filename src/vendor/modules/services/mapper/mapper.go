package mapper

import (
	"github.com/satori/go.uuid"
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
	Rooms map[string]*Room
	Tags  []string
	Fly   []*Path
}

func (m *Mapper) NewWalking() *Walking {
	walking := NewWalking()
	walking.Rooms = &m.Rooms
	walking.Tags = m.Tags
	walking.Fly = m.Fly
	return walking
}

func (m *Mapper) GetRoomID(name string) []string {
	result := []string{}
	for _, v := range m.Rooms {
		if v.Name == name {
			result = append(result, v.ID)
		}
	}
	return result
}

func (m *Mapper) GetRoomName(id string) string {
	result := m.Rooms[id]
	if result == nil {
		return ""
	}
	return result.Name
}
func (m *Mapper) AddPath(id string, p *Path) {
	room := m.Rooms[id]
	if room == nil {
		return
	}
	room.Exits = append(room.Exits, p)
}
func (m *Mapper) ClearRoom(id string) {
	room := NewRoom()
	room.ID = id
	m.Rooms[id] = room
}
func (m *Mapper) NewArea(size int) []string {
	result := []string{}
	for i := 0; i < size; i++ {
		uid, _ := uuid.NewV1()
		id := uid.String()
		result = append(result, id)
		m.ClearRoom(id)
	}
	return result
}
func (m *Mapper) GetExits(id string) []*Path {
	result := []*Path{}
	room := m.Rooms[id]
	if room == nil {
		return result
	}
	for _, v := range room.Exits {
		if len(v.Tags) == 0 && len(v.ExcludeTags) == 0 {
			result = append(result, v)
		}
	}
	return result
}

func New() *Mapper {
	return &Mapper{
		Rooms: map[string]*Room{},
		Tags:  []string{},
		Fly:   []*Path{},
	}
}
