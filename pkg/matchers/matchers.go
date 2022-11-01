package matchers

import (
	"github.com/cearius/go-charmagic/pkg/matchers/m_multibyte"
	"github.com/cearius/go-charmagic/pkg/matchers/m_singlebyte"
	"github.com/cearius/go-charmagic/pkg/matchers/m_unicode"
	"github.com/cearius/go-charmagic/pkg/matching"
)

func CreateAll() []matching.Matcher {
	return Collect(
		m_unicode.Create_Unicode_Matchers(),
		m_singlebyte.Create_SingleByte_Matchers(),
		m_multibyte.Create_MultiByte_Matchers(),
	)
}
