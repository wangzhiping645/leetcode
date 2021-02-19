package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i := -1
	flag := true
	for flag {
		i++
		if i >= len(strs[0]) {
			break
		}

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				flag = false
				break
			}
			if strs[j][i] != strs[0][i] {
				flag = false
				break
			}
		}
	}
	return strs[0][:i]
}
