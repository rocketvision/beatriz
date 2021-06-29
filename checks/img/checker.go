package img

import (
	"strings"

	"github.com/rocketvision/beatriz/core"
	"github.com/rocketvision/beatriz/util"
	"golang.org/x/net/html"
)

var formats = []string{
	".webp",
	".gif",
	".svg",
}

type Checker struct{}

func (Checker) Reset(ctx *core.Context) {}

func (Checker) EnterElement(ctx *core.Context, tag string, attrs []html.Attribute, closed bool) {
	if tag != "img" {
		return
	}

	source, hasSource := util.GetAttr(attrs, "src")
	if hasSource {
		if !checkFormat(source) {
			ctx.Issue(core.BestPractice, "img: A src deveria estar em formato WebP, GIF ou SVG.")
		}
	} else {
		ctx.Issue(core.SyntaxError, "img: Adicione o atributo src.")
	}

	hasWidth := util.HasAttr(attrs, "width")
	hasHeight := util.HasAttr(attrs, "height")
	if !hasWidth || !hasHeight {
		ctx.Issue(core.BestPractice, "img: Adicione os atributos width e height.")
	}

	if !util.HasAttr(attrs, "alt") {
		ctx.Issue(core.Accessibility, "img: Adicione o atributo alt.")
	}
}

func (Checker) LeaveElement(ctx *core.Context, tag string) {}

func checkFormat(src string) bool {
	for _, fmt := range formats {
		if strings.HasSuffix(src, fmt) {
			return true
		}
	}
	return false
}

func init() {
	core.RegisterChecker(Checker{})
}
