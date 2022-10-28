package matchers

import (
	"github.com/cearius/go-charmagic/internal/matchers/m_multibyte"
	"github.com/cearius/go-charmagic/internal/matchers/m_unicode"
	"github.com/cearius/go-charmagic/pkg/matching"
)

// var recognizers = []recognizer{
// 	newRecognizer_utf8(),
// 	newRecognizer_utf16be(),
// 	newRecognizer_utf16le(),
// 	newRecognizer_utf32be(),
// 	newRecognizer_utf32le(),
// 	newRecognizer_8859_1_en(),
// 	newRecognizer_8859_1_da(),
// 	newRecognizer_8859_1_de(),
// 	newRecognizer_8859_1_es(),
// 	newRecognizer_8859_1_fr(),
// 	newRecognizer_8859_1_it(),
// 	newRecognizer_8859_1_nl(),
// 	newRecognizer_8859_1_no(),
// 	newRecognizer_8859_1_pt(),
// 	newRecognizer_8859_1_sv(),
// 	newRecognizer_8859_2_cs(),
// 	newRecognizer_8859_2_hu(),
// 	newRecognizer_8859_2_pl(),
// 	newRecognizer_8859_2_ro(),
// 	newRecognizer_8859_5_ru(),
// 	newRecognizer_8859_6_ar(),
// 	newRecognizer_8859_7_el(),
// 	newRecognizer_8859_8_I_he(),
// 	newRecognizer_8859_8_he(),
// 	newRecognizer_windows_1251(),
// 	newRecognizer_windows_1256(),
// 	newRecognizer_KOI8_R(),
// 	newRecognizer_8859_9_tr(),

// 	newRecognizer_sjis(),
// 	newRecognizer_gb_18030(),
// 	newRecognizer_euc_jp(),
// 	newRecognizer_euc_kr(),
// 	newRecognizer_big5(),

// 	newRecognizer_2022JP(),
// 	newRecognizer_2022KR(),
// 	newRecognizer_2022CN(),

// 	newRecognizer_IBM424_he_rtl(),
// 	newRecognizer_IBM424_he_ltr(),
// 	newRecognizer_IBM420_ar_rtl(),
// 	newRecognizer_IBM420_ar_ltr(),
//}

var all = []matching.Matcher{
	m_unicode.MatchUTF8,
	m_unicode.MatchUTF16BE,
	m_unicode.MatchUTF16LE,

	m_multibyte.Match_big5(),
	m_multibyte.Match_euc_jp(),
	m_multibyte.Match_euc_kr(),
	m_multibyte.Match_gb_18030(),
}
