package world

import "strconv"

type MatchResult struct {
	List  []string
	Named map[string]string
}

func (r *MatchResult) ReplaceList(name string) []string {
	result := []string{"%%", "%", "%N", name, "%C", ""}
	for k, v := range r.List {
		if k < 9 {
			result = append(result, "@"+strconv.Itoa(k+1), v)
		} else {
			result = append(result, "@<"+strconv.Itoa(k+1)+">", v)
		}
	}
	for k, v := range r.Named {
		result = append(result, "@<"+k+">", v)
	}
	return result
}
func NewMatchResult() *MatchResult {
	return &MatchResult{
		List:  []string{},
		Named: map[string]string{},
	}
}

type Matcher interface {
	Match(message string) (*MatchResult, error)
}
