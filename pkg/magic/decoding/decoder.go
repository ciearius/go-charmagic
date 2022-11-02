package decoding

import (
	"io"

	"github.com/cearius/go-charmagic/pkg/magic/matching"
)

type DecoderCreator func() Decoder

type ByteDecoder func([]byte) []byte
type ReaderDecoder func(io.Reader) io.Reader

type Decoder interface {
	Accepts(matching.Result) bool
	DecodeBytes([]byte) ([]byte, error)
	DecodeReader(io.Reader) io.Reader
}
