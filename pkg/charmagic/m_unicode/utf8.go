package m_unicode

import (
	"bytes"
	"io"

	"github.com/cearius/go-charmagic/pkg/magic/decoding"
	"github.com/cearius/go-charmagic/pkg/magic/matching"
)

const UTF8_CharsetName = "UTF-8"

var UTF8_BOM = []byte{0xEF, 0xBB, 0xBF}

func Create_UTF8_Matcher() matching.Matcher {
	return CreateUnicodeMatcher(func(input matching.Input) (out matching.Result) {
		out.Charset = UTF8_CharsetName
		out.BOM = bytes.HasPrefix(input.Raw, UTF8_BOM)

		rawLen := len(input.Raw)

		var numValid, numInvalid uint32
		var trailBytes uint8

		for i := 0; i < rawLen; i++ {
			c := input.Raw[i]
			if c&0x80 == 0 {
				continue
			}
			if c&0xE0 == 0xC0 {
				trailBytes = 1
			} else if c&0xF0 == 0xE0 {
				trailBytes = 2
			} else if c&0xF8 == 0xF0 {
				trailBytes = 3
			} else {
				numInvalid++
				if numInvalid > 5 {
					break
				}
				trailBytes = 0
			}

			for i++; i < rawLen; i++ {
				c = input.Raw[i]
				if c&0xC0 != 0x80 {
					numInvalid++
					break
				}
				if trailBytes--; trailBytes == 0 {
					numValid++
					break
				}
			}
		}

		if out.BOM && numInvalid == 0 {
			out.Confidence = 100
		} else if out.BOM && numValid > numInvalid*10 {
			out.Confidence = 80
		} else if numValid > 3 && numInvalid == 0 {
			out.Confidence = 100
		} else if numValid > 0 && numInvalid == 0 {
			out.Confidence = 80
		} else if numValid == 0 && numInvalid == 0 {
			// Plain ASCII
			out.Confidence = 10
		} else if numValid > numInvalid*10 {
			out.Confidence = 25
		}

		return
	})
}

type utf8_Decoder struct {
}

func (d *utf8_Decoder) Accepts(t matching.Result) bool {
	return t.Charset == UTF8_CharsetName
}

func (d *utf8_Decoder) DecodeBytes(buf []byte) ([]byte, error) {
	return buf, nil
}

func (d *utf8_Decoder) DecodeReader(r io.Reader) io.Reader {
	return r
}

func Create_UTF8_Decoder() decoding.Decoder {
	return &utf8_Decoder{}
}
