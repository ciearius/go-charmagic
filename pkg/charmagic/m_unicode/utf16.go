package m_unicode

import (
	"bytes"
	"io"

	encoding "golang.org/x/text/encoding"

	"github.com/cearius/go-charmagic/pkg/magic/decoding"
	"github.com/cearius/go-charmagic/pkg/magic/matching"
	"golang.org/x/text/encoding/unicode"
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

type utf16_Decoder struct {
	big_endian bool
}

func (d *utf16_Decoder) Accepts(r matching.Result) bool {
	if d.big_endian {
		return r.Charset == UTF16BE_CharsetName
	} else {
		return r.Charset == UTF16LE_CharsetName
	}
}

func (d *utf16_Decoder) NewDecoder() *encoding.Decoder {
	var endianness unicode.Endianness

	if d.big_endian {
		endianness = unicode.BigEndian
	} else {
		endianness = unicode.LittleEndian
	}

	return unicode.UTF16(endianness, unicode.UseBOM).NewDecoder()
}

func (d *utf16_Decoder) DecodeBytes(buf []byte) ([]byte, error) {
	return d.NewDecoder().Bytes(buf)
}

func (d *utf16_Decoder) DecodeReader(r io.Reader) io.Reader {
	return d.NewDecoder().Reader(r)
}

func Create_UTF16BE_Decoder() decoding.Decoder {
	return &utf16_Decoder{big_endian: true}
}

func Create_UTF16LE_Decoder() decoding.Decoder {
	return &utf16_Decoder{big_endian: false}
}
