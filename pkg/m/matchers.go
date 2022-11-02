package m

import (
	"github.com/cearius/go-charmagic/pkg/m/multibyte"
	"github.com/cearius/go-charmagic/pkg/m/singlebyte"
	"github.com/cearius/go-charmagic/pkg/m/unicode"
	"github.com/cearius/go-charmagic/pkg/matching"
	"github.com/cearius/go-charmagic/pkg/util"
)

func CreateAllMatchers() []matching.Matcher {
	return util.Collect(
		unicode.Create_Unicode_Matchers(),
		singlebyte.Create_SingleByte_Matchers(),
		multibyte.Create_MultiByte_Matchers(),
	)
}
