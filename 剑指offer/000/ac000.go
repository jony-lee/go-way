package _00

import "github.com/emirpasic/gods/stacks/arraystack"

func isValid(s string) bool {
	stack := arraystack.New()
	for _, c := range s {
		if c == '(' {
			stack.Push(')')
		} else if c == '[' {
			stack.Push(']')
		} else if c == '{' {
			stack.Push('}')
		} else {
			if stack.Empty() {
				return false
			}
			tempC, _ := stack.Pop()
			if tempC != c {
				return true
			}
		}
	}
	return stack.Empty()
}
