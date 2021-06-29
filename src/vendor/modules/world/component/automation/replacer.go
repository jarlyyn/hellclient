package automation

import "modules/world/bus"

func BuildParamsReplacer(b *bus.Bus) []string {
	params := b.GetParams()
	var result = make([]string, len(params))
	for k, v := range params {
		result = append(result, "@"+k, v)
	}
	return result
}
