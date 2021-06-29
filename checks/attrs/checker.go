package attrs

import (
	"github.com/rocketvision/beatriz/core"
)

type Checker struct{}

func (Checker) Reset(ctx *core.Context) {}

func (Checker) EnterElement(ctx *core.Context, tag string, attrs []core.Attr, closed bool) {
	known := make(map[string]string)
	for _, attr := range attrs {
		old, isDuplicate := known[attr.Key]
		if isDuplicate {
			if attr.Val == old {
				ctx.Issue(core.SyntaxError, "%v: Atributo duplicado %v (valores idênticos).", tag, attr.Key)
			} else {
				ctx.Issue(core.SyntaxError, "%v: Atributo duplicado %v (valores diferentes).", tag, attr.Key)
			}
		}
	}
}

func (Checker) LeaveElement(ctx *core.Context, tag string) {}

func init() {
	core.RegisterChecker(Checker{})
}
