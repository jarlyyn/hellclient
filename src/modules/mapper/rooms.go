package mapper

import "github.com/herb-go/uniqueid"

type Rooms struct {
	Rooms          map[string]*Room
	TemporaryPaths map[string][]*Path
}

func NewRooms() *Rooms {
	return &Rooms{
		Rooms:          map[string]*Room{},
		TemporaryPaths: map[string][]*Path{},
	}
}
func (r *Rooms) GetRoom(id string) *Room {
	return r.Rooms[id]
}
func (r *Rooms) GetTemporaryPaths(id string) []*Path {
	return r.TemporaryPaths[id]
}
func (r *Rooms) GetRoomID(name string) []string {
	result := []string{}
	for _, v := range r.Rooms {
		if v.Name == name {
			result = append(result, v.ID)
		}
	}
	return result
}

func (r *Rooms) GetRoomName(id string) string {
	result := r.Rooms[id]
	if result == nil {
		return ""
	}
	return result.Name
}
func (r *Rooms) SetRoomName(id string, name string) {
	result := r.Rooms[id]
	if result == nil {
		result = NewRoom()
		result.ID = id
		r.Rooms[id] = result
	}
	result.Name = name

}
func (r *Rooms) AddPath(id string, p *Path) bool {
	room := r.Rooms[id]
	if room == nil {
		return false
	}
	room.Exits = append(room.Exits, p)
	return true
}
func (r *Rooms) AddTemporaryPath(id string, p *Path) bool {
	paths := r.TemporaryPaths[id]
	if paths == nil {
		paths = []*Path{}
	}
	paths = append(paths, p)
	r.TemporaryPaths[id] = paths
	return true
}

func (r *Rooms) ClearRoom(id string) {
	r.clearRoom(id)
}
func (r *Rooms) clearRoom(id string) {
	room := NewRoom()
	room.ID = id
	r.Rooms[id] = room
}

func (r *Rooms) NewArea(size int) []string {
	result := []string{}
	for i := 0; i < size; i++ {
		id := uniqueid.MustGenerateID()
		result = append(result, id)
		r.clearRoom(id)
	}
	return result
}
func (r *Rooms) GetExits(id string, tags map[string]bool, all bool) []*Path {
	result := []*Path{}
	room := r.Rooms[id]
	texits := r.TemporaryPaths[id]
	if room == nil && texits == nil {
		return result
	}
	if room != nil {
		for _, v := range room.Exits {
			if all || ValidateTags(tags, v) {
				result = append(result, v)
			}
		}
	}
	if texits != nil {
		for _, v := range texits {
			if all || ValidateTags(tags, v) {
				result = append(result, v)
			}
		}
	}
	return result
}
func (r *Rooms) Reset() {
	r.Rooms = map[string]*Room{}
	r.TemporaryPaths = map[string][]*Path{}
}
func (r *Rooms) ResetTemporary() {
	r.TemporaryPaths = map[string][]*Path{}
}
