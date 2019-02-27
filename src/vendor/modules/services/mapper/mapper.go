package mapper

import (
	"io/ioutil"
	"strconv"
	"strings"
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

type RoomAllIniLoader struct {
	TokenAfterRoomID          string
	TokenBeforeExites         string
	TokenExitsSep             string
	TokenExitsAfterTag        string
	TokenExitsAfterExcludeTag string
	TokenBeforeTarget         string
	TokenBeforeWalkLength     string
}

func (l *RoomAllIniLoader) Open(filename string) (*Mapper, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	m := New()
	lines := strings.Split(string(bs), "\n")
	for k := range lines {
		l.readData(m, lines[k])
	}
	return m, nil
}

func (l *RoomAllIniLoader) readData(m *Mapper, line string) {
	line = strings.TrimSpace(line)
	roomAndExits := strings.SplitN(line, l.TokenBeforeExites, 2)
	if len(roomAndExits) < 2 {
		return
	}
	roomIDAndName := strings.SplitN(roomAndExits[0], l.TokenAfterRoomID, 2)
	if len(roomIDAndName) < 2 {
		return
	}
	room := NewRoom()
	room.ID = strings.TrimSpace(roomIDAndName[0])
	room.Name = strings.TrimSpace(roomIDAndName[1])
	l.readExits(room, roomAndExits[1])
	m.Rooms[room.ID] = room
}
func (l *RoomAllIniLoader) readExits(r *Room, line string) {
	exits := strings.Split(line, l.TokenExitsSep)
	for k := range exits {
		p := l.exitToPath(exits[k])
		if p == nil {
			continue
		}
		p.From = r.ID
		r.Exits = append(r.Exits, p)
	}
}
func (l *RoomAllIniLoader) exitToPath(line string) *Path {
	line = strings.TrimSpace(line)
	path := NewPath()
	PathAndTarget := strings.SplitN(line, l.TokenBeforeTarget, 2)
	if len(PathAndTarget) < 2 || PathAndTarget[1] == "" {
		return nil
	}
	target := PathAndTarget[1]
	line = PathAndTarget[0]
	PathAndWalklength := strings.SplitN(target, l.TokenBeforeWalkLength, 2)
	if len(PathAndWalklength) < 2 {
		path.Delay = 1
	} else {
		var err error
		path.Delay, err = strconv.Atoi(PathAndWalklength[1])
		if err == nil {
			path.Delay = 1
		} else {
			if path.Delay < 1 {
				path.Delay = 1
			}
		}
	}
	path.To = PathAndWalklength[0]
	for _, v := range strings.SplitAfter(line, l.TokenExitsAfterTag) {
		for _, v2 := range strings.SplitAfter(v, l.TokenExitsAfterExcludeTag) {
			tag := strings.TrimSpace(v2)
			if tag == "" {
				path.Command = tag
				continue
			}
			last := tag[len(tag)-1:]
			if last == l.TokenExitsAfterTag {
				path.Tags[tag[:len(tag)-1]] = true
			} else if last == l.TokenExitsAfterExcludeTag {
				path.ExcludeTags[tag[:len(tag)-1]] = true
			} else {
				path.Command = tag
			}
		}
	}
	return path
}

var CommonRoomAllIniLoader = &RoomAllIniLoader{
	TokenAfterRoomID:          "=",
	TokenBeforeExites:         "|",
	TokenExitsSep:             ",",
	TokenExitsAfterTag:        ">",
	TokenExitsAfterExcludeTag: "<",
	TokenBeforeTarget:         ":",
	TokenBeforeWalkLength:     "%",
}

func New() *Mapper {
	return &Mapper{
		Rooms: map[string]*Room{},
		Tags:  []string{},
		Fly:   []*Path{},
	}
}
