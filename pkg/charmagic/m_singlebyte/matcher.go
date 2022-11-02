package m_singlebyte

import (
	"github.com/cearius/go-charmagic/pkg/ling"
	"github.com/cearius/go-charmagic/pkg/matching"
)

// Recognizer for single byte charset family
type SingleByteMatcher struct {
	charset        string
	HasC1_Fallback string
	language       string
	charMap        *[256]byte
	ngram          *[64]uint32
}

func Create_SingleByte_Matchers() []matching.Matcher {
	return []matching.Matcher{
		Create_IBM420_ar_rtl_Matcher(),
		Create_IBM420_ar_ltr_Matcher(),
		Create_IBM424_he_rtl_Matcher(),
		Create_IBM424_he_ltr_Matcher(),
		Create_8859_1_en_Matcher(),
		Create_8859_1_da_Matcher(),
		Create_8859_1_de_Matcher(),
		Create_8859_1_es_Matcher(),
		Create_8859_1_fr_Matcher(),
		Create_8859_1_it_Matcher(),
		Create_8859_1_nl_Matcher(),
		Create_8859_1_no_Matcher(),
		Create_8859_1_pt_Matcher(),
		Create_8859_1_sv_Matcher(),
		Create_8859_2_cs_Matcher(),
		Create_8859_2_hu_Matcher(),
		Create_8859_2_pl_Matcher(),
		Create_8859_2_ro_Matcher(),
		Create_8859_5_ru_Matcher(),
		Create_8859_6_ar_Matcher(),
		Create_8859_7_el_Matcher(),
		Create_8859_8_I_he_Matcher(),
		Create_8859_8_he_Matcher(),
		Create_8859_9_tr_Matcher(),
		Create_KOI8_R_Matcher(),
		Create_windows_1251_Matcher(),
		Create_windows_1256_Matcher(),
	}
}

func (r *SingleByteMatcher) Match(input matching.Input) matching.Result {
	res := matching.CreateResult(r.charset)

	// if the input contains C1 bytes and there is a C1 fallback defined - use it.
	if input.HasC1 && len(r.HasC1_Fallback) > 0 {
		res.Charset = r.HasC1_Fallback
	}

	res.Language = r.language
	res.Confidence = r.parseNgram(input.Raw)

	return res
}

func (r *SingleByteMatcher) parseNgram(input []byte) int {
	state := ling.NewNgramState(r.ngram)

	for _, inChar := range input {
		c := r.charMap[inChar]

		if c != 0 {
			state.AddByte(c)
		}
	}

	state.AddByte(0x20)

	rate := state.HitRate()

	if rate > 0.33 {
		return 98
	}

	return int(rate * 300)
}
