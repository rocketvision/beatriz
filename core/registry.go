package core

import (
	"golang.org/x/net/html"
)

type Checker interface {
	Reset(*Context)
	EnterElement(*Context, string, []html.Attribute, bool)
	LeaveElement(*Context, string)
}

var registry []Checker

func RegisterChecker(c Checker) {
	registry = append(registry, c)
}
