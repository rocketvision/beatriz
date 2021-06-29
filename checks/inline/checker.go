package attrs

import (
	"github.com/rocketvision/beatriz/core"
	"github.com/rocketvision/beatriz/util"
	"golang.org/x/net/html"
)

func EnterElement(ctx *core.Context, tag string, attrs []html.Attribute) {
	if util.HasAttr(attrs, "style") {
		ctx.Issue(core.BestPractice, "inline1", "%v: Evite estilização inline (style).", tag)
	}
}

func init() {
	core.RegisterChecker(&core.Checker{
		EnterElement: EnterElement,
	})
}
