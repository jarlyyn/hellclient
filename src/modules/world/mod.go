package world

type Mod struct {
	Enabled    bool
	Exists     bool
	FileList   []string
	FolderList []string
}

func NewMod() *Mod {
	return &Mod{
		Enabled:    false,
		Exists:     false,
		FileList:   []string{},
		FolderList: []string{},
	}
}
