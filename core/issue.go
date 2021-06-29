package core

type IssueType int

const (
	Generic IssueType = iota

	SyntaxError
	BestPractice
	Accessibility
)

type Issue struct {
	Type IssueType
	Pos  Position
	Text string
}
