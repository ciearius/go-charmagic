package multibyte

import (
	"errors"
	"math"

	"github.com/cearius/go-charmagic/pkg/matching"
)

var errEOB = errors.New("unexpected end of input buffer")
var errBadChar = errors.New("decoded a bad char")

type MultiByteMatcher struct {
	charset     string
	language    string
	decoder     CharDecoder
	commonChars []uint16
}

type CharDecoder interface {
	DecodeOneChar([]byte) (c uint16, remain []byte, err error)
}

func Create_MultiByte_Matchers() []matching.Matcher {
	return []matching.Matcher{
		Create_big5_Matcher(),
		Create_euc_jp_Matcher(),
		Create_euc_kr_Matcher(),
		Create_gb_18030_Matcher(),
		Create_sjis_Matcher(),
	}
}

func (r *MultiByteMatcher) Match(input matching.Input) (output matching.Result) {
	output.Charset = r.charset
	output.Language = r.language
	output.Confidence = r.matchConfidence(input)

	return
}

func (r *MultiByteMatcher) matchConfidence(input matching.Input) int {
	raw := input.Raw
	var c uint16
	var err error
	var totalCharCount, badCharCount, singleByteCharCount, doubleByteCharCount, commonCharCount int
	for c, raw, err = r.decoder.DecodeOneChar(raw); len(raw) > 0; c, raw, err = r.decoder.DecodeOneChar(raw) {
		totalCharCount++
		if err != nil {
			badCharCount++
		} else if c <= 0xFF {
			singleByteCharCount++
		} else {
			doubleByteCharCount++
			if r.commonChars != nil && binarySearch(r.commonChars, c) {
				commonCharCount++
			}
		}
		if badCharCount >= 2 && badCharCount*5 >= doubleByteCharCount {
			return 0
		}
	}

	if doubleByteCharCount <= 10 && badCharCount == 0 {
		if doubleByteCharCount == 0 && totalCharCount < 10 {
			return 0
		} else {
			return 10
		}
	}

	if doubleByteCharCount < 20*badCharCount {
		return 0
	}
	if r.commonChars == nil {
		confidence := 30 + doubleByteCharCount - 20*badCharCount
		if confidence > 100 {
			confidence = 100
		}
		return confidence
	}
	maxVal := math.Log(float64(doubleByteCharCount) / 4)
	scaleFactor := 90 / maxVal
	confidence := int(math.Log(float64(commonCharCount)+1)*scaleFactor + 10)
	if confidence > 100 {
		confidence = 100
	}
	if confidence < 0 {
		confidence = 0
	}
	return confidence
}

func binarySearch(l []uint16, c uint16) bool {
	start := 0
	end := len(l) - 1
	for start <= end {
		mid := (start + end) / 2
		if c == l[mid] {
			return true
		} else if c < l[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
