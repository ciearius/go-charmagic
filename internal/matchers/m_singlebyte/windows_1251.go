package m_singlebyte

import "github.com/cearius/go-charmagic/pkg/matching"

var charMap_windows_1251 = [256]byte{
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
	0x90, 0x83, 0x20, 0x83, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x9A, 0x20, 0x9C, 0x9D, 0x9E, 0x9F,
	0x90, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x9A, 0x20, 0x9C, 0x9D, 0x9E, 0x9F,
	0x20, 0xA2, 0xA2, 0xBC, 0x20, 0xB4, 0x20, 0x20,
	0xB8, 0x20, 0xBA, 0x20, 0x20, 0x20, 0x20, 0xBF,
	0x20, 0x20, 0xB3, 0xB3, 0xB4, 0xB5, 0x20, 0x20,
	0xB8, 0x20, 0xBA, 0x20, 0xBC, 0xBE, 0xBE, 0xBF,
	0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7,
	0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
	0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7,
	0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF,
	0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7,
	0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
	0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7,
	0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF,
}

var ngrams_windows_1251 = [64]uint32{
	0x20E220, 0x20E2EE, 0x20E4EE, 0x20E7E0, 0x20E820, 0x20EAE0, 0x20EAEE, 0x20EDE0, 0x20EDE5, 0x20EEE1, 0x20EFEE, 0x20EFF0, 0x20F0E0, 0x20F1EE, 0x20F1F2, 0x20F2EE,
	0x20F7F2, 0x20FDF2, 0xE0EDE8, 0xE0F2FC, 0xE3EE20, 0xE5EBFC, 0xE5EDE8, 0xE5F1F2, 0xE5F220, 0xE820EF, 0xE8E520, 0xE8E820, 0xE8FF20, 0xEBE5ED, 0xEBE820, 0xEBFCED,
	0xEDE020, 0xEDE520, 0xEDE8E5, 0xEDE8FF, 0xEDEE20, 0xEDEEE2, 0xEE20E2, 0xEE20EF, 0xEE20F1, 0xEEE220, 0xEEE2E0, 0xEEE3EE, 0xEEE920, 0xEEEBFC, 0xEEEC20, 0xEEF1F2,
	0xEFEEEB, 0xEFF0E5, 0xEFF0E8, 0xEFF0EE, 0xF0E0E2, 0xF0E5E4, 0xF1F2E0, 0xF1F2E2, 0xF1F2E8, 0xF1FF20, 0xF2E5EB, 0xF2EE20, 0xF2EEF0, 0xF2FC20, 0xF7F2EE, 0xFBF520,
}

func Match_windows_1251() matching.Matcher {
	return &SingleByteMatcher{
		charset:  "windows-1251",
		language: "ar",
		charMap:  &charMap_windows_1251,
		ngram:    &ngrams_windows_1251,
	}
}
