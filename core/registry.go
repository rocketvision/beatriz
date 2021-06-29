package core

import "fmt"

type Attr struct {
	Key string
	Val string
}

func (a *Attr) String() string {
	return fmt.Sprintf("%v=%q", a.Key, a.Val)
}

type Checker interface {
	Reset(*Context)
	EnterElement(*Context, string, []Attr, bool)
	LeaveElement(*Context, string)
}

var registry []Checker

func RegisterChecker(c Checker) {
	registry = append(registry, c)
}
