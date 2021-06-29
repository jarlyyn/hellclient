package world

type MatchResult struct {
	List  []string
	Named map[string]string
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
