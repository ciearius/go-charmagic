package m_singlebyte

import "github.com/cearius/go-charmagic/pkg/matching"

var charMap_8859_9 = [256]byte{
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67,
	0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F,
	0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77,
	0x78, 0x79, 0x7A, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67,
	0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F,
	0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77,
	0x78, 0x79, 0x7A, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0xAA, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0xB5, 0x20, 0x20,
	0x20, 0x20, 0xBA, 0x20, 0x20, 0x20, 0x20, 0x20,
	0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7,
	0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
	0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0x20,
	0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0x69, 0xFE, 0xDF,
	0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7,
	0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
	0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0x20,
	0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF,
}

var ngrams_8859_9_tr = [64]uint32{
	0x206261, 0x206269, 0x206275, 0x206461, 0x206465, 0x206765, 0x206861, 0x20696C, 0x206B61, 0x206B6F, 0x206D61, 0x206F6C, 0x207361, 0x207461, 0x207665, 0x207961,
	0x612062, 0x616B20, 0x616C61, 0x616D61, 0x616E20, 0x616EFD, 0x617220, 0x617261, 0x6172FD, 0x6173FD, 0x617961, 0x626972, 0x646120, 0x646520, 0x646920, 0x652062,
	0x65206B, 0x656469, 0x656E20, 0x657220, 0x657269, 0x657369, 0x696C65, 0x696E20, 0x696E69, 0x697220, 0x6C616E, 0x6C6172, 0x6C6520, 0x6C6572, 0x6E2061, 0x6E2062,
	0x6E206B, 0x6E6461, 0x6E6465, 0x6E6520, 0x6E6920, 0x6E696E, 0x6EFD20, 0x72696E, 0x72FD6E, 0x766520, 0x796120, 0x796F72, 0xFD6E20, 0xFD6E64, 0xFD6EFD, 0xFDF0FD,
}

func Match_8859_9(language string, ngram *[64]uint32) matching.Matcher {
	return &SingleByteMatcher{
		charset:          "ISO-8859-9",
		hasC1ByteCharset: "windows-1254",
		language:         language,
		charMap:          &charMap_8859_9,
		ngram:            ngram,
	}
}

func Match_8859_9_tr() matching.Matcher {
	return Match_8859_9("tr", &ngrams_8859_9_tr)
}
