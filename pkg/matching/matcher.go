package matching

type Matcher interface {
	Match(input MatcherInput) MatchResult
}
