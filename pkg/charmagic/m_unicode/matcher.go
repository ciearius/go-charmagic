package m_unicode

import "github.com/cearius/go-charmagic/pkg/magic/matching"

type UnicodeMatcher struct {
	mf func(mi matching.Input) matching.Result
}

func (m *UnicodeMatcher) Match(i matching.Input) matching.Result {
	return m.mf(i)
}

func CreateUnicodeMatcher(f func(mi matching.Input) matching.Result) *UnicodeMatcher {
	return &UnicodeMatcher{f}
}

func Create_Unicode_Matchers() []matching.Matcher {
	return []matching.Matcher{
		Create_UTF8_Matcher(),
		Create_UTF16LE_Matcher(),
		Create_UTF16BE_Matcher(),
		// TODO: add utf32 matcher
	}
}
