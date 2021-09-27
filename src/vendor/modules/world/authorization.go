package world

type Authorization struct {
	Callback string
	Items    []string
	Reason   string
}

func CreateAuthorization(callback string, items []string, reason string) *Authorization {
	return &Authorization{
		Callback: callback,
		Items:    append([]string{}, items...),
		Reason:   reason,
	}
}
