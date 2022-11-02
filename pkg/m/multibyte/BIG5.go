package multibyte

type charDecoder_big5 struct {
}

func (charDecoder_big5) DecodeOneChar(input []byte) (c uint16, remain []byte, err error) {
	if len(input) == 0 {
		return 0, nil, errEOB
	}
	first := input[0]
	remain = input[1:]
	c = uint16(first)
	if first <= 0x7F || first == 0xFF {
		return
	}
	if len(remain) == 0 {
		return c, nil, errEOB
	}
	second := remain[0]
	remain = remain[1:]
	c = c<<8 | uint16(second)
	if second < 0x40 || second == 0x7F || second == 0xFF {
		err = errBadChar
	}
	return
}

var commonChars_big5 = []uint16{
	0xa140, 0xa141, 0xa142, 0xa143, 0xa147, 0xa149, 0xa175, 0xa176, 0xa440, 0xa446,
	0xa447, 0xa448, 0xa451, 0xa454, 0xa457, 0xa464, 0xa46a, 0xa46c, 0xa477, 0xa4a3,
	0xa4a4, 0xa4a7, 0xa4c1, 0xa4ce, 0xa4d1, 0xa4df, 0xa4e8, 0xa4fd, 0xa540, 0xa548,
	0xa558, 0xa569, 0xa5cd, 0xa5e7, 0xa657, 0xa661, 0xa662, 0xa668, 0xa670, 0xa6a8,
	0xa6b3, 0xa6b9, 0xa6d3, 0xa6db, 0xa6e6, 0xa6f2, 0xa740, 0xa751, 0xa759, 0xa7da,
	0xa8a3, 0xa8a5, 0xa8ad, 0xa8d1, 0xa8d3, 0xa8e4, 0xa8fc, 0xa9c0, 0xa9d2, 0xa9f3,
	0xaa6b, 0xaaba, 0xaabe, 0xaacc, 0xaafc, 0xac47, 0xac4f, 0xacb0, 0xacd2, 0xad59,
	0xaec9, 0xafe0, 0xb0ea, 0xb16f, 0xb2b3, 0xb2c4, 0xb36f, 0xb44c, 0xb44e, 0xb54c,
	0xb5a5, 0xb5bd, 0xb5d0, 0xb5d8, 0xb671, 0xb7ed, 0xb867, 0xb944, 0xbad8, 0xbb44,
	0xbba1, 0xbdd1, 0xc2c4, 0xc3b9, 0xc440, 0xc45f,
}

func Create_big5_Matcher() *MultiByteMatcher {
	return &MultiByteMatcher{
		"Big5",
		"zh",
		charDecoder_big5{},
		commonChars_big5,
	}
}