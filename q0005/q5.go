package main

import "fmt"

func main() {
	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("cbbd"))
	//fmt.Println(longestPalindrome("a"))
	//fmt.Println(longestPalindrome("ac"))
	//fmt.Println(longestPalindrome("acaaa"))
	//fmt.Println(longestPalindrome("bb"))
	//fmt.Println(longestPalindrome("aaaa"))
	//fmt.Println(longestPalindrome("aaaaa"))
	fmt.Println(longestPalindrome("abbcccbbbcaaccbababcbcabca"))

}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	l := len(s)
	result := s[:1]
	max := 1
	for i := 0; i < l; i++ {
		flag1 := true
		flag2 := true
		for j := 1; i+j < l && (flag1 || flag2); j++ {
			if flag1 && i-j >= 0 && s[i-j] == s[i+j] {
				if 2*j+1 > max {
					max = 2*j + 1
					result = s[i-j : i+j+1]
				}
			} else {
				flag1 = false
			}

			if !flag2 {
				continue
			}

			//fmt.Println(i)
			if i+1 >= l || s[i] != s[i+1] {
				flag2 = false
				continue
			}
			if 2 > max {
				max = 2
				result = s[i : i+2]
			}

			if i+j+1 >= l || i-j < 0 || s[i-j] != s[i+j+1] {
				flag2 = false
				continue
			}
			if 2*j+2 > max {
				max = 2*j + 2
				result = s[i-j : i+j+2]
			}
		}
	}
	return result
}

func longestPalindromDP(s string) string {
	if s == "" {
		return ""
	}
	result := s[:1]
	max := 1
	l := len(s)
	dp := [2][]bool{}
	for i := 0; i < 2; i++ {
		dp[i] = make([]bool, l)
	}
	for i := l - 1; i >= 0; i-- {
		dp[0] = dp[1]
		dp[1] = make([]bool, l)
		for j := i; j < l; j++ {
			//dp[0][j] = dp[1][j]
			//dp[1][j] = false
			if i == j {
				dp[1][j] = true
				continue
			}
			if j == i+1 && s[i] == s[j] {
				dp[1][j] = true
				if 2 > max {
					max = 2
					result = s[i : j+1]
				}
				continue
			}
			if dp[0][j-1] && s[i] == s[j] {
				dp[1][j] = true
				if j-i+1 > max {
					//fmt.Println(i, j, j-i+1, dp[0][j-1], s[i:j+1])
					max = j - i + 1
					result = s[i : j+1]
				}
			}
		}
	}

	return result
}

func longestPalindrome1(s string) string {
	if s == "" {
		return ""
	}
	result := s[:1]
	max := 1
	l := len(s)
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = true
		if i < l-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
			if 2 > max {
				max = 2
				result = s[i : i+2]
			}
		}
	}

	for i := l - 3; i >= 0; i-- {
		for j := l - 1; j > i; j-- {
			if s[i] == s[j] && dp[i+1][j-1] {
				//fmt.Println(i, j)
				dp[i][j] = true
				if j-i+1 > max {
					max = j - i + 1
					result = s[i : j+1]
				}
			}
		}
	}

	return result
}

func longestPalindromeBest(s string) string {
	subS := [2]int{0, 0}

	for i := 0; i < len(s); i++ {
		// 回文字符串的长度为奇数
		k := 1
		for i-k >= 0 && i+k < len(s) && s[i-k] == s[i+k] {
			if 2*k+1 > subS[1]-subS[0] {
				subS[0] = i - k
				subS[1] = i + k
			}
			k++
		}

		if i+1 < len(s) && s[i] == s[i+1] {
			k = 0
			//  回文字符串的长度为偶数
			for i-k >= 0 && i+k+1 < len(s) && s[i-k] == s[i+k+1] {
				if 2*k+2 > subS[1]-subS[0] {
					subS[0] = i - k
					subS[1] = i + k + 1
				}
				k++
			}

		}

	}
	return s[subS[0] : subS[1]+1]
}
