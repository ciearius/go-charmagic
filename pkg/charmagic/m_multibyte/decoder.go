package m_multibyte

import (
	"io"

	"github.com/cearius/go-charmagic/pkg/magic/decoding"
	"github.com/cearius/go-charmagic/pkg/magic/matching"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

type MultiByteDecoder struct {
	name string
	enc  encoding.Encoding
}

func (d *MultiByteDecoder) Accepts(t matching.Result) bool {
	return t.Charset == d.name
}

func (d *MultiByteDecoder) DecodeBytes(buf []byte) ([]byte, error) {
	return d.enc.NewDecoder().Bytes(buf)
}

func (d *MultiByteDecoder) DecodeReader(r io.Reader) io.Reader {
	return d.enc.NewDecoder().Reader(r)
}

func Create_MultiByteDecoder_Decoders() []decoding.Decoder {
	decoders := []decoding.Decoder{}

	encodings := make(map[string]encoding.Encoding)

	encodings["Big5"] = traditionalchinese.Big5
	encodings["EUC-JP"] = japanese.EUCJP
	encodings["EUC-KR"] = korean.EUCKR
	encodings["GB-18030"] = simplifiedchinese.GB18030
	encodings["Shift_JIS"] = japanese.ShiftJIS

	for name, enc := range encodings {
		decoders = append(decoders, &MultiByteDecoder{name, enc})
	}

	return decoders
}
