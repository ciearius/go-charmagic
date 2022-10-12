package matching

type MatchResult struct {
	Charset    string
	Language   string
	Confidence int
	BOM        bool
}

type MatchResults []MatchResult

func (mr MatchResults) Len() int {
	return len(mr)
}

func (ps MatchResults) Less(i, j int) bool {
	return ps[i].Confidence < ps[j].Confidence
}

// Swap switches two values in our slice
func (mr MatchResults) Swap(i, j int) {
	tmp := mr[i]

	mr[i] = mr[j]

	mr[j] = tmp
}
