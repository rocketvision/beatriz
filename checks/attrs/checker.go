package attrs

import (
	"github.com/rocketvision/beatriz/core"
	"golang.org/x/net/html"
)

func EnterElement(ctx *core.Context, tag string, attrs []html.Attribute) {
	known := make(map[string]string)
	for _, attr := range attrs {
		old, isDuplicate := known[attr.Key]
		if isDuplicate {
			if attr.Val == old {
				ctx.Issue(core.SyntaxError, "attr1", "%v: Atributo duplicado %v (valores idênticos).", tag, attr.Key)
			} else {
				ctx.Issue(core.SyntaxError, "attr2", "%v: Atributo duplicado %v (valores diferentes).", tag, attr.Key)
			}
		}
	}
}

func init() {
	core.RegisterChecker(&core.Checker{
		EnterElement: EnterElement,
	})
}
