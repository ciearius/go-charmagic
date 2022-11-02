package matching

type Results []Result

// Len is returning
func (mr Results) Len() int {
	return len(mr)
}

// Less determines if the first element's value is less than the second's
func (ps Results) Less(i, j int) bool {
	return ps[i].Confidence < ps[j].Confidence
}

// Swap switches two values in slice
func (mr Results) Swap(i, j int) {
	tmp := mr[i]

	mr[i] = mr[j]

	mr[j] = tmp
}
