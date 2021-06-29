package stack

import (
	"github.com/rocketvision/beatriz/core"
	"golang.org/x/net/html"
)

func Reset(ctx *core.Context) {
	stack = nil
	broken = false
}

func EnterElement(ctx *core.Context, tag string, attrs []html.Attribute) {
	if broken && !Persistent {
		return
	}
	if NoCloseTags[tag] {
		return
	}
	pushTag(tag)
}

func LeaveElement(ctx *core.Context, tag string) {
	if NoCloseTags[tag] {
		ctx.Issue(core.SyntaxError, "stack1", "%v: Não feche esse elemento.", tag)
		return
	}

	if broken && !Persistent {
		return
	}
	if stack == nil {
		ctx.Issue(core.SyntaxError, "stack2", "%v: Elemento fechado na raiz do documento.", tag)
		// c.broken = true
		return
	}
	innermost := peekTag()
	if tag != innermost {
		ctx.Issue(core.SyntaxError, "stack3", "%v: Feche o elemento %v anterior.", tag, innermost)
		broken = true
		return
	}
	popTag()
}

func init() {
	core.RegisterChecker(&core.Checker{
		EnterElement: EnterElement,
		LeaveElement: LeaveElement,
	})
}
