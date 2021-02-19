package main

import "strings"

func main() {

}

var (
	m = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
)

func intToRoman1(num int) string {
	s := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	for _, v := range s {
		if num < v {
			continue
		}
		res += strings.Repeat(m[v], num/v)
		num = num % v
	}
	return res
}

func intToRoman(num int) string {
	if r, found := m[num]; found {
		return r
	}
	if num > 1000 {
		return "M" + intToRoman(num-1000)
	} else if num > 900 {
		return "CM" + intToRoman(num-900)
	} else if num > 500 {
		return "D" + intToRoman(num-500)
	} else if num > 400 {
		return "CD" + intToRoman(num-400)
	} else if num > 100 {
		return "C" + intToRoman(num-100)
	} else if num > 90 {
		return "XC" + intToRoman(num-90)
	} else if num > 50 {
		return "L" + intToRoman(num-50)
	} else if num > 40 {
		return "XL" + intToRoman(num-40)
	} else if num > 10 {
		return "X" + intToRoman(num-10)
	} else if num > 9 {
		return "IX" + intToRoman(num-9)
	} else if num > 5 {
		return "V" + intToRoman(num-5)
	} else if num > 4 {
		return "Iv" + intToRoman(num-4)
	} else {
		return "I" + intToRoman(num-1)
	}
}
