package piscine

import (
	"fmt"
	"os"
)

func Brackets() {
	if len(os.Args) < 2 {
		return
	}

	for i := 1; i < len(os.Args); i++ {
		if isBalanced(os.Args[i]) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}

func isBalanced(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if isOpenBracket(ch) {
			stack = append(stack, ch)
		} else if isCloseBracket(ch) {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if !matches(top, ch) {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isOpenBracket(ch rune) bool {
	return ch == '(' || ch == '[' || ch == '{'
}

func isCloseBracket(ch rune) bool {
	return ch == ')' || ch == ']' || ch == '}'
}

func matches(open, close rune) bool {
	return (open == '(' && close == ')') ||
		(open == '[' && close == ']') ||
		(open == '{' && close == '}')
}
