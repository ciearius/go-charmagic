package m_singlebyte

import "github.com/cearius/go-charmagic/pkg/matching"

var charMap_8859_1 = [256]byte{
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
	0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xDF,
	0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7,
	0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
	0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0x20,
	0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF,
}

var ngrams_8859_1_en = [64]uint32{
	0x206120, 0x20616E, 0x206265, 0x20636F, 0x20666F, 0x206861, 0x206865, 0x20696E, 0x206D61, 0x206F66, 0x207072, 0x207265, 0x207361, 0x207374, 0x207468, 0x20746F,
	0x207768, 0x616964, 0x616C20, 0x616E20, 0x616E64, 0x617320, 0x617420, 0x617465, 0x617469, 0x642061, 0x642074, 0x652061, 0x652073, 0x652074, 0x656420, 0x656E74,
	0x657220, 0x657320, 0x666F72, 0x686174, 0x686520, 0x686572, 0x696420, 0x696E20, 0x696E67, 0x696F6E, 0x697320, 0x6E2061, 0x6E2074, 0x6E6420, 0x6E6720, 0x6E7420,
	0x6F6620, 0x6F6E20, 0x6F7220, 0x726520, 0x727320, 0x732061, 0x732074, 0x736169, 0x737420, 0x742074, 0x746572, 0x746861, 0x746865, 0x74696F, 0x746F20, 0x747320,
}

var ngrams_8859_1_da = [64]uint32{
	0x206166, 0x206174, 0x206465, 0x20656E, 0x206572, 0x20666F, 0x206861, 0x206920, 0x206D65, 0x206F67, 0x2070E5, 0x207369, 0x207374, 0x207469, 0x207669, 0x616620,
	0x616E20, 0x616E64, 0x617220, 0x617420, 0x646520, 0x64656E, 0x646572, 0x646574, 0x652073, 0x656420, 0x656465, 0x656E20, 0x656E64, 0x657220, 0x657265, 0x657320,
	0x657420, 0x666F72, 0x676520, 0x67656E, 0x676572, 0x696765, 0x696C20, 0x696E67, 0x6B6520, 0x6B6B65, 0x6C6572, 0x6C6967, 0x6C6C65, 0x6D6564, 0x6E6465, 0x6E6520,
	0x6E6720, 0x6E6765, 0x6F6720, 0x6F6D20, 0x6F7220, 0x70E520, 0x722064, 0x722065, 0x722073, 0x726520, 0x737465, 0x742073, 0x746520, 0x746572, 0x74696C, 0x766572,
}

var ngrams_8859_1_de = [64]uint32{
	0x20616E, 0x206175, 0x206265, 0x206461, 0x206465, 0x206469, 0x206569, 0x206765, 0x206861, 0x20696E, 0x206D69, 0x207363, 0x207365, 0x20756E, 0x207665, 0x20766F,
	0x207765, 0x207A75, 0x626572, 0x636820, 0x636865, 0x636874, 0x646173, 0x64656E, 0x646572, 0x646965, 0x652064, 0x652073, 0x65696E, 0x656974, 0x656E20, 0x657220,
	0x657320, 0x67656E, 0x68656E, 0x687420, 0x696368, 0x696520, 0x696E20, 0x696E65, 0x697420, 0x6C6963, 0x6C6C65, 0x6E2061, 0x6E2064, 0x6E2073, 0x6E6420, 0x6E6465,
	0x6E6520, 0x6E6720, 0x6E6765, 0x6E7465, 0x722064, 0x726465, 0x726569, 0x736368, 0x737465, 0x742064, 0x746520, 0x74656E, 0x746572, 0x756E64, 0x756E67, 0x766572,
}

var ngrams_8859_1_es = [64]uint32{
	0x206120, 0x206361, 0x20636F, 0x206465, 0x20656C, 0x20656E, 0x206573, 0x20696E, 0x206C61, 0x206C6F, 0x207061, 0x20706F, 0x207072, 0x207175, 0x207265, 0x207365,
	0x20756E, 0x207920, 0x612063, 0x612064, 0x612065, 0x61206C, 0x612070, 0x616369, 0x61646F, 0x616C20, 0x617220, 0x617320, 0x6369F3, 0x636F6E, 0x646520, 0x64656C,
	0x646F20, 0x652064, 0x652065, 0x65206C, 0x656C20, 0x656E20, 0x656E74, 0x657320, 0x657374, 0x69656E, 0x69F36E, 0x6C6120, 0x6C6F73, 0x6E2065, 0x6E7465, 0x6F2064,
	0x6F2065, 0x6F6E20, 0x6F7220, 0x6F7320, 0x706172, 0x717565, 0x726120, 0x726573, 0x732064, 0x732065, 0x732070, 0x736520, 0x746520, 0x746F20, 0x756520, 0xF36E20,
}

var ngrams_8859_1_fr = [64]uint32{
	0x206175, 0x20636F, 0x206461, 0x206465, 0x206475, 0x20656E, 0x206574, 0x206C61, 0x206C65, 0x207061, 0x20706F, 0x207072, 0x207175, 0x207365, 0x20736F, 0x20756E,
	0x20E020, 0x616E74, 0x617469, 0x636520, 0x636F6E, 0x646520, 0x646573, 0x647520, 0x652061, 0x652063, 0x652064, 0x652065, 0x65206C, 0x652070, 0x652073, 0x656E20,
	0x656E74, 0x657220, 0x657320, 0x657420, 0x657572, 0x696F6E, 0x697320, 0x697420, 0x6C6120, 0x6C6520, 0x6C6573, 0x6D656E, 0x6E2064, 0x6E6520, 0x6E7320, 0x6E7420,
	0x6F6E20, 0x6F6E74, 0x6F7572, 0x717565, 0x72206C, 0x726520, 0x732061, 0x732064, 0x732065, 0x73206C, 0x732070, 0x742064, 0x746520, 0x74696F, 0x756520, 0x757220,
}

var ngrams_8859_1_it = [64]uint32{
	0x20616C, 0x206368, 0x20636F, 0x206465, 0x206469, 0x206520, 0x20696C, 0x20696E, 0x206C61, 0x207065, 0x207072, 0x20756E, 0x612063, 0x612064, 0x612070, 0x612073,
	0x61746F, 0x636865, 0x636F6E, 0x64656C, 0x646920, 0x652061, 0x652063, 0x652064, 0x652069, 0x65206C, 0x652070, 0x652073, 0x656C20, 0x656C6C, 0x656E74, 0x657220,
	0x686520, 0x692061, 0x692063, 0x692064, 0x692073, 0x696120, 0x696C20, 0x696E20, 0x696F6E, 0x6C6120, 0x6C6520, 0x6C6920, 0x6C6C61, 0x6E6520, 0x6E6920, 0x6E6F20,
	0x6E7465, 0x6F2061, 0x6F2064, 0x6F2069, 0x6F2073, 0x6F6E20, 0x6F6E65, 0x706572, 0x726120, 0x726520, 0x736920, 0x746120, 0x746520, 0x746920, 0x746F20, 0x7A696F,
}

var ngrams_8859_1_nl = [64]uint32{
	0x20616C, 0x206265, 0x206461, 0x206465, 0x206469, 0x206565, 0x20656E, 0x206765, 0x206865, 0x20696E, 0x206D61, 0x206D65, 0x206F70, 0x207465, 0x207661, 0x207665,
	0x20766F, 0x207765, 0x207A69, 0x61616E, 0x616172, 0x616E20, 0x616E64, 0x617220, 0x617420, 0x636874, 0x646520, 0x64656E, 0x646572, 0x652062, 0x652076, 0x65656E,
	0x656572, 0x656E20, 0x657220, 0x657273, 0x657420, 0x67656E, 0x686574, 0x696520, 0x696E20, 0x696E67, 0x697320, 0x6E2062, 0x6E2064, 0x6E2065, 0x6E2068, 0x6E206F,
	0x6E2076, 0x6E6465, 0x6E6720, 0x6F6E64, 0x6F6F72, 0x6F7020, 0x6F7220, 0x736368, 0x737465, 0x742064, 0x746520, 0x74656E, 0x746572, 0x76616E, 0x766572, 0x766F6F,
}

var ngrams_8859_1_no = [64]uint32{
	0x206174, 0x206176, 0x206465, 0x20656E, 0x206572, 0x20666F, 0x206861, 0x206920, 0x206D65, 0x206F67, 0x2070E5, 0x207365, 0x20736B, 0x20736F, 0x207374, 0x207469,
	0x207669, 0x20E520, 0x616E64, 0x617220, 0x617420, 0x646520, 0x64656E, 0x646574, 0x652073, 0x656420, 0x656E20, 0x656E65, 0x657220, 0x657265, 0x657420, 0x657474,
	0x666F72, 0x67656E, 0x696B6B, 0x696C20, 0x696E67, 0x6B6520, 0x6B6B65, 0x6C6520, 0x6C6C65, 0x6D6564, 0x6D656E, 0x6E2073, 0x6E6520, 0x6E6720, 0x6E6765, 0x6E6E65,
	0x6F6720, 0x6F6D20, 0x6F7220, 0x70E520, 0x722073, 0x726520, 0x736F6D, 0x737465, 0x742073, 0x746520, 0x74656E, 0x746572, 0x74696C, 0x747420, 0x747465, 0x766572,
}

var ngrams_8859_1_pt = [64]uint32{
	0x206120, 0x20636F, 0x206461, 0x206465, 0x20646F, 0x206520, 0x206573, 0x206D61, 0x206E6F, 0x206F20, 0x207061, 0x20706F, 0x207072, 0x207175, 0x207265, 0x207365,
	0x20756D, 0x612061, 0x612063, 0x612064, 0x612070, 0x616465, 0x61646F, 0x616C20, 0x617220, 0x617261, 0x617320, 0x636F6D, 0x636F6E, 0x646120, 0x646520, 0x646F20,
	0x646F73, 0x652061, 0x652064, 0x656D20, 0x656E74, 0x657320, 0x657374, 0x696120, 0x696361, 0x6D656E, 0x6E7465, 0x6E746F, 0x6F2061, 0x6F2063, 0x6F2064, 0x6F2065,
	0x6F2070, 0x6F7320, 0x706172, 0x717565, 0x726120, 0x726573, 0x732061, 0x732064, 0x732065, 0x732070, 0x737461, 0x746520, 0x746F20, 0x756520, 0xE36F20, 0xE7E36F,
}

var ngrams_8859_1_sv = [64]uint32{
	0x206174, 0x206176, 0x206465, 0x20656E, 0x2066F6, 0x206861, 0x206920, 0x20696E, 0x206B6F, 0x206D65, 0x206F63, 0x2070E5, 0x20736B, 0x20736F, 0x207374, 0x207469,
	0x207661, 0x207669, 0x20E472, 0x616465, 0x616E20, 0x616E64, 0x617220, 0x617474, 0x636820, 0x646520, 0x64656E, 0x646572, 0x646574, 0x656420, 0x656E20, 0x657220,
	0x657420, 0x66F672, 0x67656E, 0x696C6C, 0x696E67, 0x6B6120, 0x6C6C20, 0x6D6564, 0x6E2073, 0x6E6120, 0x6E6465, 0x6E6720, 0x6E6765, 0x6E696E, 0x6F6368, 0x6F6D20,
	0x6F6E20, 0x70E520, 0x722061, 0x722073, 0x726120, 0x736B61, 0x736F6D, 0x742073, 0x746120, 0x746520, 0x746572, 0x74696C, 0x747420, 0x766172, 0xE47220, 0xF67220,
}

func Match_8859_1(language string, ngram *[64]uint32) matching.Matcher {
	return &SingleByteMatcher{
		charset:          "ISO-8859-1",
		hasC1ByteCharset: "windows-1252",
		language:         language,
		charMap:          &charMap_8859_1,
		ngram:            ngram,
	}
}

func Match_8859_1_en() matching.Matcher {
	return Match_8859_1("en", &ngrams_8859_1_en)
}
func Match_8859_1_da() matching.Matcher {
	return Match_8859_1("da", &ngrams_8859_1_da)
}
func Match_8859_1_de() matching.Matcher {
	return Match_8859_1("de", &ngrams_8859_1_de)
}
func Match_8859_1_es() matching.Matcher {
	return Match_8859_1("es", &ngrams_8859_1_es)
}
func Match_8859_1_fr() matching.Matcher {
	return Match_8859_1("fr", &ngrams_8859_1_fr)
}
func Match_8859_1_it() matching.Matcher {
	return Match_8859_1("it", &ngrams_8859_1_it)
}
func Match_8859_1_nl() matching.Matcher {
	return Match_8859_1("nl", &ngrams_8859_1_nl)
}
func Match_8859_1_no() matching.Matcher {
	return Match_8859_1("no", &ngrams_8859_1_no)
}
func Match_8859_1_pt() matching.Matcher {
	return Match_8859_1("pt", &ngrams_8859_1_pt)
}
func Match_8859_1_sv() matching.Matcher {
	return Match_8859_1("sv", &ngrams_8859_1_sv)
}
