package detecting

import (
	"github.com/cearius/go-charmagic/internal/matchers/m_unicode"
	"github.com/cearius/go-charmagic/pkg/matching"
)

var matchers = []matching.Matcher{
	m_unicode.MatchUTF8,
	m_unicode.MatchUTF16BE,
	m_unicode.MatchUTF16LE,
}
