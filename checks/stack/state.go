package stack

var (
	stack  []string
	broken bool
)

func pushTag(tag string) {
	stack = append(stack, tag)
}

func popTag() {
	stack = stack[:len(stack)-1]
}

func peekTag() string {
	return stack[len(stack)-1]
}
