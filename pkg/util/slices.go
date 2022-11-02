package util

// Collect appends all slices together
func Collect[K any](ks ...[]K) []K {
	res := []K{}

	for _, k := range ks {
		res = append(res, k...)
	}

	return res
}
