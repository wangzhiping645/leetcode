package main

import "fmt"

func main() {
	fmt.Println(convert2("PAYPALISHIRING", 3))
	fmt.Println(convert2("PAYPALISHIRING", 4))
}

func convert2(s string, numRows int) string {
	l := len(s)
	if numRows == 1 || l <= 1 {
		return s
	}
	rows := numRows
	if rows > l {
		rows = l
	}

	result := make([]byte, 0, l)
	max := l/(2*rows-2) + 1
	for i := 0; i < rows; i++ {
		for j := 0; j < max; j++ {
			k := 2*j*(rows-1) + i
			if k >= l {
				break
			}
			result = append(result, s[k])
			//result += s[k : k+1]
			if i == 0 || i == rows-1 {
				continue
			}
			k = k + (rows-i-1)*2
			if k >= l {
				break
			}
			result = append(result, s[k])
		}
	}

	return string(result)
}

func convert1(s string, numRows int) string {
	if s == "" || numRows == 1 {
		return s
	}
	//l := len(s)
	rows := make([]string, numRows)
	//for i := 0; i < numRows; i++ {
	//	rows[i] = make(string, 0, l/numRows+1)
	//}

	row := 0
	next := 1
	for i := 0; i < len(s); i++ {
		if row == 0 {
			next = 1
		} else if row == numRows-1 {
			next = -1
		}
		rows[row] += string(s[i])
		row += next
	}

	result := ""
	for _, row := range rows {
		result += string(row)
	}
	//fmt.Println(rows)
	//fmt.Println(len(result))
	//fmt.Println("11", result, "22")
	return result
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	l := len(s)
	rows := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		rows[i] = make([]byte, 0, l/numRows+1)
	}

	row := 0
	next := 1
	for i := 0; i < len(s); i++ {
		if row == 0 {
			next = 1
		} else if row == numRows-1 {
			next = -1
		}
		rows[row] = append(rows[row], s[i])
		row += next
	}

	result := ""
	for _, row := range rows {
		result += string(row)
	}
	//fmt.Println(rows)
	//fmt.Println(len(result))
	//fmt.Println("11", result, "22")
	return result
}
