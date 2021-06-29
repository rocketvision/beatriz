package a

import (
	"github.com/rocketvision/beatriz/core"
	"github.com/rocketvision/beatriz/util"
	"golang.org/x/net/html"
)

func EnterElement(ctx *core.Context, tag string, attrs []html.Attribute) {
	if tag != "a" {
		return
	}
	href, hasHref := util.GetAttr(attrs, "href")
	if !hasHref {
		ctx.Issue(core.SyntaxError, "a1", "a: Adicione o atributo href.")
		return
	}
	if InvalidHREF[href] {
		ctx.Issue(core.Accessibility, "a2", "a: O atributo href deve conter uma URL válida.")
	}
}

func init() {
	core.RegisterChecker(&core.Checker{
		EnterElement: EnterElement,
	})
}
