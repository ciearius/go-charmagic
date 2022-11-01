package ling

type NgramState struct {
	ngram                uint32
	ignoreSpace          bool
	ngramCount, ngramHit uint32
	table                *[64]uint32
}

func NewNgramState(table *[64]uint32) *NgramState {
	return &NgramState{
		ngram:       0,
		ignoreSpace: false,
		ngramCount:  0,
		ngramHit:    0,
		table:       table,
	}
}

func (s *NgramState) AddByte(b byte) {
	const ngramMask = 0xFFFFFF
	if !(b == 0x20 && s.ignoreSpace) {
		s.ngram = ((s.ngram << 8) | uint32(b)) & ngramMask
		s.ignoreSpace = (s.ngram == 0x20)
		s.ngramCount++
		if s.lookup() {
			s.ngramHit++
		}
	}
	s.ignoreSpace = (b == 0x20)
}

func (s *NgramState) HitRate() float32 {
	if s.ngramCount == 0 {
		return 0
	}
	return float32(s.ngramHit) / float32(s.ngramCount)
}

func (s *NgramState) lookup() bool {
	var index int
	if s.table[index+32] <= s.ngram {
		index += 32
	}
	if s.table[index+16] <= s.ngram {
		index += 16
	}
	if s.table[index+8] <= s.ngram {
		index += 8
	}
	if s.table[index+4] <= s.ngram {
		index += 4
	}
	if s.table[index+2] <= s.ngram {
		index += 2
	}
	if s.table[index+1] <= s.ngram {
		index += 1
	}
	if s.table[index] > s.ngram {
		index -= 1
	}
	if index < 0 || s.table[index] != s.ngram {
		return false
	}
	return true
}
