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
	Code string
	Text string
	Pos  Position
}
