package world

type ClientInfo struct {
	ID       string
	ReadyAt  int64
	HostPort string
	ScriptID string
	Running  bool
}

type ClientInfos []*ClientInfo

// Len is the number of elements in the collection.
func (info ClientInfos) Len() int {
	return len(info)
}

// Less reports whether the element with index i
func (info ClientInfos) Less(i, j int) bool {
	return info[i].ReadyAt > info[j].ReadyAt
}

// Swap swaps the elements with indexes i and j.
func (info ClientInfos) Swap(i, j int) {
	info[i], info[j] = info[j], info[i]
}
