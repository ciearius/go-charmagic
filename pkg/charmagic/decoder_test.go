package charmagic_test

import (
	"testing"

	"github.com/cearius/go-charmagic/pkg/charmagic"
	"github.com/cearius/go-charmagic/pkg/matching"
	testify "github.com/stretchr/testify/assert"
)

func Test_AllDecoders(t *testing.T) {
	assert := testify.New(t)
	ti := matching.FromBytes([]byte{})

	for _, d := range charmagic.CreateAllMatchers() {
		r := d.Match(ti)

		_, err := charmagic.GetDecoder(r.Charset)

		assert.NoError(err, r.Charset+" encoding is missing")
	}
}

func Test_AllDecoders_AllSupported(t *testing.T) {
	assert := testify.New(t)
	ti := matching.FromBytes([]byte{})

	for _, d := range charmagic.CreateAllMatchers() {
		r := d.Match(ti)

		assert.NotPanics(func() {
			enc := charmagic.GetDecoderFromResult(&r)
			assert.NotNil(enc)
		})
	}
}
