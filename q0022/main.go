package main

import "fmt"

func main() {
	fmt.Println(len(generateParenthesis(1)))
	fmt.Println(len(generateParenthesis(2)))
	fmt.Println(len(generateParenthesis(3)))
	fmt.Println(len(generateParenthesis(4)))
	fmt.Println(len(generateParenthesis(5)))
}

var (
	result    = []string{}
	singleMax = 0
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return result
	}
	singleMax = n
	result = make([]string, 0)
	backtrack("(", 1, 0)
	return result
}

func backtrack(before string, leftNum, rightNum int) {

	if leftNum < singleMax {
		backtrack(before+"(", leftNum+1, rightNum)
	}
	if rightNum < leftNum {
		backtrack(before+")", leftNum, rightNum+1)
	}
	if leftNum == rightNum && rightNum == singleMax {
		result = append(result, before)
	}
}
