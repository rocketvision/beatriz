package core

import "fmt"

type Context struct {
	state *State
	pos   Position
}

func NewContext(state *State) *Context {
	return &Context{
		state: state,
		pos: Position{
			Line:   1,
			Column: 1,
		},
	}
}

func (c *Context) Issue(typ IssueType, code string, format string, args ...interface{}) {
	c.state.issues = append(c.state.issues, Issue{
		Type: typ,
		Code: code,
		Text: fmt.Sprintf(format, args...),
		Pos:  c.pos,
	})
}

// func (c *Context) IssueError(typ IssueType, err error) {
// 	c.Issue(typ, err.Error())
// }

func (c *Context) Issues() []Issue {
	return c.state.issues
}
