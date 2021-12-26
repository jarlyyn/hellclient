package version

import (
	"fmt"

	"github.com/herb-go/misc/version"
)

const Major = 0
const Year = 2021
const Month = 11
const Day = 29
const Patch = 0
const Build = ""

const Message = "Hellclient version %s\n"

var Version = &version.DateVersion{
	Major: Major,
	Year:  Year,
	Month: Month,
	Day:   Day,
	Patch: Patch,
	Build: Build,
}

func init() {
	fmt.Printf(Message, Version.FullVersionCode())
}