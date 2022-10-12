package m_singlebyte

import "github.com/cearius/go-charmagic/pkg/matching"

// Recognizer for single byte charset family
type SingleByteMatcher struct {
	charset          string
	hasC1ByteCharset string
	language         string
	charMap          *[256]byte
	ngram            *[64]uint32
}

func (r *SingleByteMatcher) Match(input matching.MatcherInput) (out matching.MatchResult) {
	out.Charset = r.charset

	if input.IsC1 && len(r.hasC1ByteCharset) > 0 {
		out.Charset = r.hasC1ByteCharset
	}

	out.Language = r.language
	out.Confidence = r.parseNgram(input.Raw)

	return
}

func (r *SingleByteMatcher) parseNgram(input []byte) int {
	state := newNgramState(r.ngram)
	for _, inChar := range input {
		c := r.charMap[inChar]
		if c != 0 {
			state.AddByte(c)
		}
	}
	state.AddByte(0x20)
	rate := state.HitRate()
	if rate > 0.33 {
		return 98
	}
	return int(rate * 300)
}
