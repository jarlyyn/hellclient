package world

type Timer struct {
	Name                  string
	Enabled               bool
	Hour                  int
	Minute                int
	Second                int
	Send                  string
	ScriptName            string
	AtTime                bool
	SendTo                int
	ActionWhenDisconnectd bool
	Temporary             bool
	OneShot               bool
	SpeedWalk             bool
	Group                 string
	OmitFromLog           bool
	OmitFromOutput        bool
}
