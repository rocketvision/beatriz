package attrs

import (
	"github.com/rocketvision/beatriz/core"
	"github.com/rocketvision/beatriz/util"
	"golang.org/x/net/html"
)

type Checker struct{}

func (Checker) Reset(ctx *core.Context) {}

func (Checker) EnterElement(ctx *core.Context, tag string, attrs []html.Attribute, closed bool) {
	if util.HasAttr(attrs, "style") {
		ctx.Issue(core.BestPractice, "%v: Evite estilização inline (style).", tag)
	}
}

func (Checker) LeaveElement(ctx *core.Context, tag string) {}

func init() {
	core.RegisterChecker(Checker{})
}
