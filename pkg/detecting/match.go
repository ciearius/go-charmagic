package detecting

import (
	"errors"
	"sort"

	"github.com/cearius/go-charmagic/pkg/charmagic"
	"github.com/cearius/go-charmagic/pkg/matching"
)

var ErrNoMatchFound = errors.New("no match found")

// MatchAll matches all supported encodings against the input bytes and assigns a confidence score.
// The higher the score the more likely the algorithm deemed the chance of having picked the right encoding.
func MatchAll(buf []byte) (results matching.Results) {
	input := matching.FromBytes(buf)

	for _, m := range charmagic.CreateAllMatchers() {
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
