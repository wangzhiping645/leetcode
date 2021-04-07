package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("aaaaa", "aa"))
	fmt.Println(strStr("a", "a"))

}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	result := -1
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if j == len(needle) {
			result = i
			break
		}
	}
	return result
}
