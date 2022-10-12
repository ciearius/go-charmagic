package m_unicode

import (
	"bytes"

	"github.com/cearius/go-charmagic/pkg/matching"
)

var (
	UTF16BE_BOM = []byte{0xFE, 0xFF}
	UTF16LE_BOM = []byte{0xFF, 0xFE}
	// utf32beBom = []byte{0x00, 0x00, 0xFE, 0xFF}
	// utf32leBom = []byte{0xFF, 0xFE, 0x00, 0x00}
)

const (
	UTF16BE_CharsetName = "UTF-16BE"
	UTF16LE_CharsetName = "UTF-16LE"
)

var MatchUTF16BE = matching.CreateMatcher(func(mi matching.MatcherInput) (out matching.MatchResult) {
	out.Charset = UTF16BE_CharsetName

	if bytes.HasPrefix(mi.Raw, UTF16BE_BOM) {
		out.Confidence = 100
	}

	return
})

var MatchUTF16LE = matching.CreateMatcher(func(mi matching.MatcherInput) (out matching.MatchResult) {
	out.Charset = UTF16LE_CharsetName

	if bytes.HasPrefix(mi.Raw, UTF16LE_BOM) {
		out.Confidence = 100
	}

	return
})
