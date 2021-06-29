package basic

import (
	"github.com/rocketvision/beatriz/core"
	"golang.org/x/net/html"
)

func Reset(ctx *core.Context) {
	sawDoctype = false
}

func EnterElement(ctx *core.Context, tag string, attrs []html.Attribute) {
	if tag == "html" {
		if !sawDoctype {
			ctx.Issue(core.SyntaxError, "basic1", "Adicione o Doctype no começo do arquivo.")
		}
	}
}

func EnterDoctype(ctx *core.Context, text string) {
	sawDoctype = true
}

func init() {
	core.RegisterChecker(&core.Checker{
		Reset:        Reset,
		EnterElement: EnterElement,
		EnterDoctype: EnterDoctype,
	})
}
