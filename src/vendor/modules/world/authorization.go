package world

type Authorization struct {
	World  string
	Script string
	Items  []string
	Reason string
}

func CreateAuthorization(world string, items []string, reason string, script string) *Authorization {
	return &Authorization{
		World:  world,
		Script: script,
		Items:  append([]string{}, items...),
		Reason: reason,
	}
}
