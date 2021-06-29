package core

import "golang.org/x/net/html"

type Checker struct {
	Reset             func(*Context)
	EnterElement      func(*Context, string, []html.Attribute)
	LeaveElement      func(*Context, string)
	EnterLeaveElement func(*Context, string, []html.Attribute)
	EnterComment      func(*Context, string)
	EnterDoctype      func(*Context, string)
	EnterText         func(*Context, string)
}

func (c *Checker) callReset(ctx *Context) {
	if c.Reset != nil {
		c.Reset(ctx)
	}
}

func (c *Checker) callEnterElement(ctx *Context, tag string, attrs []html.Attribute) {
	if c.EnterElement != nil {
		c.EnterElement(ctx, tag, attrs)
	}
}

func (c *Checker) callLeaveElement(ctx *Context, tag string) {
	if c.LeaveElement != nil {
		c.LeaveElement(ctx, tag)
	}
}

func (c *Checker) callEnterLeaveElement(ctx *Context, tag string, attrs []html.Attribute) {
	if c.EnterLeaveElement != nil {
		c.EnterLeaveElement(ctx, tag, attrs)
	} else {
		c.callEnterElement(ctx, tag, attrs)
		c.callLeaveElement(ctx, tag)
	}
}

func (c *Checker) callEnterComment(ctx *Context, text string) {
	if c.EnterComment != nil {
		c.EnterComment(ctx, text)
	}
}

func (c *Checker) callEnterDoctype(ctx *Context, text string) {
	if c.EnterDoctype != nil {
		c.EnterDoctype(ctx, text)
	}
}

func (c *Checker) callEnterText(ctx *Context, text string) {
	if c.EnterText != nil {
		c.EnterText(ctx, text)
	}
}
