package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("({[]})"))
}

func isRight(c byte) bool {
	return c == ']' || c == '}' || c == ')'
}

func isLeft(c byte) bool {
	return c == '[' || c == '{' || c == '('
}

func getLeft(c byte) byte {
	if c == ']' {
		return '['
	}

	if c == '}' {
		return '{'
	}

	if c == ')' {
		return '('
	}
	return '0'
}

func isValid(s string) bool {
	l := len(s)
	if l%2 == 1 {
		return false
	}
	stack := make([]byte, l/2)
	index := -1
	for i := 0; i < l; i++ {
		t := s[i]
		if isLeft(t) {
			index++
			if index >= l/2 || index < 0 {
				return false
			}
			stack[index] = t
			continue
		}

		if index >= l/2 || index < 0 {
			return false
		}
		if stack[index] == getLeft(t) {
			index--
		} else {
			return false
		}

	}
	if index != -1 {
		return false
	}

	return true
}
