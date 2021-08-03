package automation

import "modules/world/bus"

func BuildParamsReplacer(b *bus.Bus) []string {
	params := b.GetParams()
	var result = make([]string, 0, len(params)+4)
	result = append(result, "\\\\", "\\", "\\@", "@")
	for k, v := range params {
		result = append(result, "@"+k, v)
	}
	return result
}
