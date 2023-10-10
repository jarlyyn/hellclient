package hud

import (
	"modules/world"
	"modules/world/bus"
	"sync"
)

const MaxSize = 20

type HUD struct {
	Locker  sync.Mutex
	content []*world.Line
}

func New() *HUD {
	return &HUD{}
}
func (h *HUD) SetSize(bus *bus.Bus, size int) {
	h.Locker.Lock()
	defer h.Locker.Unlock()
	if size < 0 {
		size = 0
	}
	if size > MaxSize {
		size = MaxSize
	}
	h.content = make([]*world.Line, size)
	bus.RaiseHUDContentEvent(append([]*world.Line{}, h.content...))
}

func (h *HUD) GetSize() int {
	h.Locker.Lock()
	defer h.Locker.Unlock()
	return len(h.content)
}
func (h *HUD) UpdateContent(bus *bus.Bus, start int, lines []*world.Line) bool {
	h.Locker.Lock()
	defer h.Locker.Unlock()
	if start < 0 {
		return false
	}
	if (start + len(lines) - 1) > len(h.content) {
		return false
	}
	for k := range lines {
		h.content[start+k] = lines[k]
	}
	bus.RaiseHUDUpdateEvent(world.CreateDiffLines(start, lines))
	return true
}
func (h *HUD) GetContent() []*world.Line {
	h.Locker.Lock()
	defer h.Locker.Unlock()
	return append([]*world.Line{}, h.content...)
}

func (h *HUD) InstallTo(b *bus.Bus) {
	b.GetHUDSize = h.GetSize
	b.SetHUDSize = b.WrapHandleInt(h.SetSize)
	b.GetHUDContent = h.GetContent
	b.UpdateHUDContent = func(start int, content []*world.Line) bool {
		return h.UpdateContent(b, start, content)
	}
}
