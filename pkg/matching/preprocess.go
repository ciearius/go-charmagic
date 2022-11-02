package matching

func FromBytes(buf []byte) Input {
	om := computeOccurances(buf)
	isC1 := isControlCodePresent(om)

	return Input{
		Raw:              buf,
		OccurranceMatrix: om,
		HasC1:            isC1,
	}
}

func computeOccurances(buf []byte) []int {
	r := make([]int, 256)
	for _, c := range buf {
		r[c] += 1
	}
	return r
}

func isControlCodePresent(om []int) bool {
	for _, count := range om[0x80 : 0x9F+1] {
		if count > 0 {
			return true
		}
	}

	return false
}
