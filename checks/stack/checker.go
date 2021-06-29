package stack

import (
	"github.com/rocketvision/beatriz/core"
)

var ignore = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}

type Checker struct {
	stack      []string
	broken     bool
	persistent bool
}

func (c *Checker) Reset(ctx *core.Context) {
	c.stack = nil
	c.broken = false

	c.persistent = true // TODO: Debug-only.
}

func (c *Checker) EnterElement(ctx *core.Context, tag string, attrs []core.Attr, closed bool) {
	if c.broken && !c.persistent {
		return
	}
	if ignore[tag] {
		return
	}
	c.stack = append(c.stack, tag)
}

func (c *Checker) LeaveElement(ctx *core.Context, tag string) {
	if ignore[tag] {
		ctx.Issue(core.SyntaxError, "%v: Não feche esse elemento.", tag)
		return
	}

	if c.broken && !c.persistent {
		return
	}
	if c.stack == nil {
		ctx.Issue(core.SyntaxError, "%v: Elemento fechado na raiz do documento.", tag)
		// c.broken = true
		return
	}
	innermost := c.stack[len(c.stack)-1]
	if tag != innermost {
		ctx.Issue(core.SyntaxError, "%v: Feche o elemento %v anterior.", tag, innermost)
		c.broken = true
		return
	}
	c.stack = c.stack[:len(c.stack)-1]
}

func init() {
	core.RegisterChecker(&Checker{})
}
