package matching

type Matcher interface {
	Match(input Input) Result
}

type MatcherCreator func() Matcher
