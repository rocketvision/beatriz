package core

type State struct {
	issues []Issue
}

func NewState() *State {
	return &State{}
}
