package unicode

import (
	"bytes"

	"github.com/cearius/go-charmagic/pkg/matching"
)

var (
	UTF16_BE_BOM = []byte{0xFE, 0xFF}
	UTF16_LE_BOM = []byte{0xFF, 0xFE}
)

const (
	UTF16BE_CharsetName = "UTF-16BE"
	UTF16LE_CharsetName = "UTF-16LE"
)

func Create_UTF16BE_Matcher() matching.Matcher {
	return CreateUnicodeMatcher(func(mi matching.Input) matching.Result {
		res := matching.CreateResult(
			UTF16BE_CharsetName,
			matching.WithBOM(bytes.HasPrefix(mi.Raw, UTF16_BE_BOM)),
		)

		if res.BOM {
			res.Confidence = 100
		}

		return res
	})
}

func Create_UTF16LE_Matcher() matching.Matcher {
	return CreateUnicodeMatcher(func(mi matching.Input) matching.Result {
		res := matching.CreateResult(UTF16LE_CharsetName)

		if bytes.HasPrefix(mi.Raw, UTF16_LE_BOM) {
			res.Confidence = 100
		}

		return res
	})
}
