package img

import (
	"strings"

	"github.com/rocketvision/beatriz/core"
	"github.com/rocketvision/beatriz/util"
	"golang.org/x/net/html"
)

func EnterElement(ctx *core.Context, tag string, attrs []html.Attribute) {
	if tag != "img" {
		return
	}

	source, hasSource := util.GetAttr(attrs, "src")
	if hasSource {
		if !acceptFormat(source) {
			ctx.Issue(core.BestPractice, "img1", "img: A src deveria estar em formato WebP, GIF ou SVG.")
		}
	} else {
		ctx.Issue(core.SyntaxError, "img2", "img: Adicione o atributo src.")
	}

	hasWidth := util.HasAttr(attrs, "width")
	hasHeight := util.HasAttr(attrs, "height")
	if !hasWidth || !hasHeight {
		ctx.Issue(core.BestPractice, "img3", "img: Adicione os atributos width e height.")
	}

	if !util.HasAttr(attrs, "alt") {
		ctx.Issue(core.Accessibility, "img4", "img: Adicione o atributo alt.")
	}
}

func acceptFormat(src string) bool {
	for _, fmt := range AcceptedFormats {
		if strings.HasSuffix(src, fmt) {
			return true
		}
	}
	return false
}

func init() {
	core.RegisterChecker(&core.Checker{
		EnterElement: EnterElement,
	})
}
