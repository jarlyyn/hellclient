package world

import "strconv"

const StringYes = "y"

func ToStringBool(v bool) string {
	if v {
		return StringYes
	}
	return ""
}

func FromStringBool(v string) bool {
	return v == StringYes || v == "1"
}
func FromStringInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		i = 0
	}
	return i
}
func FromStringFloat(v string) float64 {
	i, err := strconv.ParseFloat(v, 64)
	if err != nil {
		i = 0
	}
	return i
}

const UpdateOK = 0
const UpdateFailNotFound = 1
const UpdateFailDuplicateName = 2
