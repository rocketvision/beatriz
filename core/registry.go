package core

var registry []*Checker

func RegisterChecker(c *Checker) {
	registry = append(registry, c)
}
