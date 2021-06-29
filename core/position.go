package core

import "fmt"

type Position struct {
	Line   int
	Column int
}

func (p Position) String() string {
	return fmt.Sprintf("Line %v, Col %v", p.Line, p.Column)
}
