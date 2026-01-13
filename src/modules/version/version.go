package version

import (
	"github.com/herb-go/misc/version"
)

const Major = 1
const Year = 2026
const Month = 1
const Day = 13
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
