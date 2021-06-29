package a

import (
	"github.com/rocketvision/beatriz/core"
	"github.com/rocketvision/beatriz/util"
)

var invalid = map[string]bool{
	"#":                   true,
	"javascript:void(0)":  true,
	"javascript:void(0);": true,
}

type Checker struct{}

func (Checker) Reset(ctx *core.Context) {}

func (Checker) EnterElement(ctx *core.Context, tag string, attrs []core.Attr, closed bool) {
	if tag != "a" {
		return
	}
	href, hasHref := util.GetAttr(attrs, "href")
	if !hasHref {
		ctx.Issue(core.SyntaxError, "a: Adicione o atributo href.")
		return
	}
	if invalid[href] {
		ctx.Issue(core.Accessibility, "a: O atributo href deve conter uma URL válida.")
	}
}

func (Checker) LeaveElement(ctx *core.Context, tag string) {}

func init() {
	core.RegisterChecker(Checker{})
}
