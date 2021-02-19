package main

import "fmt"

func main() {
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-91283472332"))
	//fmt.Println(myAtoi("   42"))
}

// z^31 = 2147483648

func myAtoi(s string) int {

	if s == "" {
		return 0
	}

	// 找第一个正确字符
	negativeFlag := -1
	needStep2 := false
	i := 0
	l := len(s)
	for ; i < l; i++ {
		t := s[i]
		if t == 32 { // 空格继续遍历
			continue
		} else if t == 48 { // 0
			negativeFlag = 1
			needStep2 = true
			break
		} else if t >= 49 && t <= 57 { // 正整数
			negativeFlag = 1
			break
		} else if t == 43 { // +
			i++
			negativeFlag = 1
			needStep2 = true
			break
		} else if t == 45 { // -
			i++
			needStep2 = true
			break
		} else {
			return 0
		}
	}

	// 遍历到第一个非0字符
	if needStep2 {
		for ; i < l; i++ {
			t := s[i]
			if t == 48 {
				continue
			} else if t >= 49 || t <= 57 { // 正整数
				break
			} else {
				return 0
			}
		}
	}
	result := 0
	for ; i < l; i++ {
		t := s[i]
		if t < 48 || t > 57 {
			break
		}

		result = result*10 + negativeFlag*int(t-48)
		//if result >= 1<<31-1 {
		//	if !negativeFlag { // 正数
		//		return 1<<31 - 1
		//	} else { // 负数
		//		if result > 1<<31-1 {
		//			return -1 << 31
		//		}
		//	}
		//	break
		//}
		if result < -1<<31 {
			result = -1 << 31
		} else if result > 1<<31-1 {
			result = 1<<31 - 1
		}
	}

	if result < -1<<31 {
		result = -1 << 31
	} else if result > 1<<31-1 {
		result = 1<<31 - 1
	}

	return result
}
