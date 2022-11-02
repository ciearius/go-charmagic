package m_singlebyte

import "github.com/cearius/go-charmagic/pkg/matching"

var charMap_8859_8 = [256]byte{
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
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0xB5, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7,
	0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
	0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7,
	0xF8, 0xF9, 0xFA, 0x20, 0x20, 0x20, 0x20, 0x20,
}

var ngrams_8859_8_I_he = [64]uint32{
	0x20E0E5, 0x20E0E7, 0x20E0E9, 0x20E0FA, 0x20E1E9, 0x20E1EE, 0x20E4E0, 0x20E4E5, 0x20E4E9, 0x20E4EE, 0x20E4F2, 0x20E4F9, 0x20E4FA, 0x20ECE0, 0x20ECE4, 0x20EEE0,
	0x20F2EC, 0x20F9EC, 0xE0FA20, 0xE420E0, 0xE420E1, 0xE420E4, 0xE420EC, 0xE420EE, 0xE420F9, 0xE4E5E0, 0xE5E020, 0xE5ED20, 0xE5EF20, 0xE5F820, 0xE5FA20, 0xE920E4,
	0xE9E420, 0xE9E5FA, 0xE9E9ED, 0xE9ED20, 0xE9EF20, 0xE9F820, 0xE9FA20, 0xEC20E0, 0xEC20E4, 0xECE020, 0xECE420, 0xED20E0, 0xED20E1, 0xED20E4, 0xED20EC, 0xED20EE,
	0xED20F9, 0xEEE420, 0xEF20E4, 0xF0E420, 0xF0E920, 0xF0E9ED, 0xF2EC20, 0xF820E4, 0xF8E9ED, 0xF9EC20, 0xFA20E0, 0xFA20E1, 0xFA20E4, 0xFA20EC, 0xFA20EE, 0xFA20F9,
}

var ngrams_8859_8_he = [64]uint32{
	0x20E0E5, 0x20E0EC, 0x20E4E9, 0x20E4EC, 0x20E4EE, 0x20E4F0, 0x20E9F0, 0x20ECF2, 0x20ECF9, 0x20EDE5, 0x20EDE9, 0x20EFE5, 0x20EFE9, 0x20F8E5, 0x20F8E9, 0x20FAE0,
	0x20FAE5, 0x20FAE9, 0xE020E4, 0xE020EC, 0xE020ED, 0xE020FA, 0xE0E420, 0xE0E5E4, 0xE0EC20, 0xE0EE20, 0xE120E4, 0xE120ED, 0xE120FA, 0xE420E4, 0xE420E9, 0xE420EC,
	0xE420ED, 0xE420EF, 0xE420F8, 0xE420FA, 0xE4EC20, 0xE5E020, 0xE5E420, 0xE7E020, 0xE9E020, 0xE9E120, 0xE9E420, 0xEC20E4, 0xEC20ED, 0xEC20FA, 0xECF220, 0xECF920,
	0xEDE9E9, 0xEDE9F0, 0xEDE9F8, 0xEE20E4, 0xEE20ED, 0xEE20FA, 0xEEE120, 0xEEE420, 0xF2E420, 0xF920E4, 0xF920ED, 0xF920FA, 0xF9E420, 0xFAE020, 0xFAE420, 0xFAE5E9,
}

func Create_8859_8_Matcher(language string, ngram *[64]uint32) *SingleByteMatcher {
	return &SingleByteMatcher{
		charset:        "ISO-8859-8",
		HasC1_Fallback: "windows-1255",
		language:       language,
		charMap:        &charMap_8859_8,
		ngram:          ngram,
	}
}

func Create_8859_8_I_he_Matcher() matching.Matcher {
	r := Create_8859_8_Matcher("he", &ngrams_8859_8_I_he)
	r.charset = "ISO-8859-8-I"
	return r
}

func Create_8859_8_he_Matcher() matching.Matcher {
	return Create_8859_8_Matcher("he", &ngrams_8859_8_he)
}
