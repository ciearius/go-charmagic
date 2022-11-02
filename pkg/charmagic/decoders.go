package charmagic

import (
	"github.com/cearius/go-charmagic/pkg/charmagic/m_multibyte"
	"github.com/cearius/go-charmagic/pkg/charmagic/m_unicode"
	"github.com/cearius/go-charmagic/pkg/magic/decoding"
	"github.com/cearius/go-charmagic/pkg/util"
)

func CreateAllDecoders() []decoding.Decoder {
	return util.Collect(
		m_unicode.Create_Unicode_Decoders(),
		m_multibyte.Create_MultiByteDecoder_Decoders(),
	)
}
