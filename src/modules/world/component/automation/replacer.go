package automation

import (
	"modules/world/bus"
	"regexp"
	"sort"
)

type Paramkeys []string

// Len is the number of elements in the collection.
func (k Paramkeys) Len() int {
	return len(k)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (k Paramkeys) Less(i, j int) bool {
	return k[i] < k[j]
}

// Swap swaps the elements with indexes i and j.
func (k Paramkeys) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func BuildParamsReplacer(b *bus.Bus) []string {
	params := b.GetParams()
	var result = make([]string, 0, len(params)+4)
	result = append(result, "\\\\", "\\", "\\@", "@")
	keys := make(Paramkeys, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(keys))
	for _, key := range keys {
		result = append(result, "@"+key, params[key])
	}
	return result
}
func BuildParamsTriggerReplacer(b *bus.Bus) []string {
	params := b.GetParams()
	var result = make([]string, 0, len(params)+4)
	result = append(result, "\\\\", "\\", "\\@", "@")
	keys := make(Paramkeys, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(keys))
	for _, key := range keys {
		result = append(result, "@"+key, params[key], "@!"+key, regexp.QuoteMeta(params[key]))
	}
	return result
}
