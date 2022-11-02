package m_unicode

import (
	"github.com/cearius/go-charmagic/pkg/magic/decoding"
)

func Create_Unicode_Decoders() []decoding.Decoder {
	return []decoding.Decoder{
		Create_UTF8_Decoder(),
		Create_UTF16BE_Decoder(),
		Create_UTF16LE_Decoder(),
	}
}
