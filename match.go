package main

import (
	"errors"
	"sort"

	"github.com/cearius/go-charmagic/pkg/m/multibyte"
	"github.com/cearius/go-charmagic/pkg/m/singlebyte"
	"github.com/cearius/go-charmagic/pkg/m/unicode"
	"github.com/cearius/go-charmagic/pkg/matching"
	"github.com/cearius/go-charmagic/pkg/util"
)

var ErrNoMatchFound = errors.New("no match found")

func CreateAllMatchers() []matching.Matcher {
	return util.Collect(
		unicode.Create_Unicode_Matchers(),
		singlebyte.Create_SingleByte_Matchers(),
		multibyte.Create_MultiByte_Matchers(),
	)
}

// MatchAll matches all supported encodings against the input bytes and assigns a confidence score.
// The higher the score the more likely the algorithm deemed the chance of having picked the right encoding.
func MatchAll(buf []byte) (results matching.Results) {
	input := matching.FromBytes(buf)

	for _, m := range CreateAllMatchers() {
		results = append(results, m.Match(input))
	}

	sort.Stable(results)

	return results
}

// MatchBest is a shorthand for selecting the highest confidence result.
func MatchBest(buf []byte) (matching.Result, error) {
	results := MatchAll(buf)

	if len(results) >= 1 {
		return results[0], nil
	}

	return matching.Result{}, ErrNoMatchFound
}
