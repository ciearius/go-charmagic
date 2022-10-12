package detecting

import (
	"errors"
	"sort"

	"github.com/cearius/go-charmagic/internal/input"
	"github.com/cearius/go-charmagic/pkg/matching"
)

var ErrNoMatchFound = errors.New("no match found")

// DetectAll matches all supported encodings against the input bytes and assigns a confidence score.
// The higher the score the more likely the algorithm deemed the chance of having picked the right encoding.
func DetectAll(buf []byte) (results matching.MatchResults) {
	input := input.FromBytes(buf)

	for _, m := range matchers {
		results = append(results, m.Match(input))
	}

	sort.Stable(results)

	return results

}

// DetectBest is a shorthand for selecting the highest confidence result found by DetectAll.
func DetectBest(buf []byte) (matching.MatchResult, error) {

	results := DetectAll(buf)

	if len(results) >= 1 {
		return results[0], nil
	}

	return matching.MatchResult{}, ErrNoMatchFound

}
