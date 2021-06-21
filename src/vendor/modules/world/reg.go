package world

import "regexp"

var IDRegexp = regexp.MustCompile(`^[0-9a-zA-Z\-\_\@\.\[\]\(\)\+]*$`)
