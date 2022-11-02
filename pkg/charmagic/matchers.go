package charmagic

import (
	"github.com/cearius/go-charmagic/pkg/charmagic/m_multibyte"
	"github.com/cearius/go-charmagic/pkg/charmagic/m_singlebyte"
	"github.com/cearius/go-charmagic/pkg/charmagic/m_unicode"
	"github.com/cearius/go-charmagic/pkg/magic/matching"
	"github.com/cearius/go-charmagic/pkg/util"
)

func CreateAll() []matching.Matcher {
	return util.Collect(
		m_unicode.Create_Unicode_Matchers(),
		m_singlebyte.Create_SingleByte_Matchers(),
		m_multibyte.Create_MultiByte_Matchers(),
	)
}
