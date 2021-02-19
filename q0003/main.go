package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
	fmt.Println()
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println()
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println()
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println()
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println()
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring1(s string) int {
	l := len(s)
	i, j, result := 0, 0, 0
	m := map[uint8]struct{}{}
	for ; j < l; j++ {
		tmp := s[j]
		_, found := m[tmp]
		if found {
			for ; i <= j-1; i++ {
				//fmt.Println(i, j)
				delete(m, s[i])
				if s[i] == tmp {
					break
				}
			}
			i++
		}
		if j-i+1 > result {
			result = j - i + 1
		}

		m[s[j]] = struct{}{}
	}
	return result
}

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	i, j, result := 0, -1, 0
	var freq [256]int
	for i < l {
		if j+1 < l && freq[s[j+1]] == 0 {
			j++
			freq[s[j]]++
		} else {
			freq[s[i]]--
			i++
		}
		if j-i+1 > result {
			result = j - i + 1
		}
	}

	return result
}
