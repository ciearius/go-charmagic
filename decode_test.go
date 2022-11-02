package main_test

import (
	"testing"

	charmagic "github.com/cearius/go-charmagic"
	"github.com/cearius/go-charmagic/pkg/m"
	"github.com/cearius/go-charmagic/pkg/matching"
	testify "github.com/stretchr/testify/assert"
)

func Test_AllDecoders(t *testing.T) {
	assert := testify.New(t)
	ti := matching.FromBytes([]byte{})

	for _, d := range m.CreateAllMatchers() {
		r := d.Match(ti)

		_, err := charmagic.GetDecoder(r.Charset)

		assert.NoError(err, r.Charset+" encoding is missing")
	}
}

// FIXME: this test will fail since not all encodings are supported!
func Test_AllDecoders_AllSupported(t *testing.T) {
	assert := testify.New(t)
	ti := matching.FromBytes([]byte{})

	for _, d := range m.CreateAllMatchers() {
		r := d.Match(ti)

		assert.NotPanics(func() {
			enc := charmagic.GetDecoderFromResult(&r)
			assert.NotNil(enc)
		})
	}
}
