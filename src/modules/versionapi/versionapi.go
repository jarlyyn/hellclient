package versionapi

import (
	appversion "modules/version"

	"fmt"

	"github.com/herb-go/misc/version"
)

const Major = 2
const Year = 2026
const Month = 8
const Day = 31
const Patch = 0
const Build = ""

const Message = "Hellclient version %s (API %s)\n"

var Version = &version.DateVersion{
	Major: Major,
	Year:  Year,
	Month: Month,
	Day:   Day,
	Patch: Patch,
	Build: Build,
}

func init() {
	fmt.Printf(Message, appversion.Version.FullVersionCode(), Version.FullVersionCode())
}
