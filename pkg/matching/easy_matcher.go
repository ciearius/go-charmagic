package matching

type EasyMatcher struct {
	matchFunc func(MatcherInput) MatchResult
}

func CreateMatcher(m func(MatcherInput) MatchResult) Matcher {
	return &EasyMatcher{m}
}

func (e EasyMatcher) Match(i MatcherInput) MatchResult {
	return e.matchFunc(i)
}
