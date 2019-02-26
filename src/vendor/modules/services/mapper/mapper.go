package mapper

import (
	"container/list"
	"io/ioutil"
	"strings"
)

type Path struct {
	content     string
	delay       string
	from        string
	to          string
	tags        map[string]bool
	excludeTags map[string]bool
}

func NewPath() *Path {
	return &Path{
		tags:        map[string]bool{},
		excludeTags: map[string]bool{},
	}
}

type Room struct {
	ID       string
	Name     string
	Exits    *list.List
	TagExits *list.List
}

func NewRoom() *Room {
	return &Room{
		Exits:    list.New(),
		TagExits: list.New(),
	}
}

type Mapper struct {
	Rooms map[string]*Room
}

func (m *Mapper) Clean() {
	m.Rooms = map[string]*Room{}
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
	}
}
func (l *RoomAllIniLoader) exitToPath(line string) *Path {
	line = strings.TrimSpace(line)
	path := NewPath()
	PathAndTarget := strings.Split(line, l.TokenBeforeTarget)
	if len(PathAndTarget) < 2 || PathAndTarget[1] == "" {
		return nil
	}
	path.to = PathAndTarget[1]
	line = PathAndTarget[0]
	for _, v := range strings.SplitAfter(line, l.TokenExitsAfterTag) {
		for _, v2 := range strings.SplitAfter(v, l.TokenExitsAfterExcludeTag) {
			tag := strings.TrimSpace(v2)
			if tag == "" {
				path.content = tag
				continue
			}
			last := tag[len(tag)-1:]
			if last == l.TokenExitsAfterTag {
				path.tags[tag[:len(tag)-1]] = true
			} else if last == l.TokenExitsAfterExcludeTag {
				path.excludeTags[tag[:len(tag)-1]] = true
			} else {
				path.content = tag
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
	TokenBeforeTarget:         ">",
	TokenBeforeWalkLength:     "%",
}

func New() *Mapper {
	return &Mapper{
		Rooms: map[string]*Room{},
	}
}
