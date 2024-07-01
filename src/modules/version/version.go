package version

import (
	"github.com/herb-go/misc/version"
)

const Major = 1
const Year = 2024
const Month = 07
const Day = 01
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
