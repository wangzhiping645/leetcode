package main

import "fmt"

func main() {
	fmt.Print(letterCombinations("233"))
}

var (
	result = []string{}
	m      = map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
)

// 回溯
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	l := 1
	for i := 0; i < len(digits); i++ {
		l *= len(m[digits[i]])
	}
	result = make([]string, 0, l)

	backtrack(&digits, -1, "")
	return result
}

func backtrack(digits *string, index int, before string) {
	index++
	if index == len(*digits)-1 {
		for _, letter := range m[(*digits)[index]] {
			result = append(result, before+string(letter))
		}
		return
	}
	for _, letter := range m[(*digits)[index]] {
		backtrack(digits, index, before+string(letter))
	}
}
