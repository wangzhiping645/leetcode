package main

var (
	m = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
)

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += m[string(s[i])]
			continue
		}
		t2 := s[i : i+2]
		v, found := m[t2]
		if found {
			result += v
			i++
		} else {
			result += m[s[i:i+1]]
		}

	}
	return result
}

var (
	m2 = map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

func romanToInt2(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += m2[s[i]]
			continue
		}
		if m2[s[i]] < m2[s[i+1]] {
			result -= m2[s[i]]
		} else {
			result += m2[s[i]]
		}

	}
	return result
}
