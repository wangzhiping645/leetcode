package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
		//want1 string
	}{
		{
			"abccc", args{s: "abccc"}, "ccc",
		},
		{
			"babcc", args{s: "babcc"}, "bab",
		},
		{
			"cbbd", args{s: "cbbd"}, "bb",
		},
		{
			"a", args{s: "a"}, "a",
		},
		{
			"ac", args{s: "ac"}, "a",
		},
		{
			"acaaa", args{s: "acaaa"}, "aaa,aca",
		},
		{
			"aaaa", args{"aaaa"}, "aaaa",
		},
		{
			"bb", args{"bb"}, "bb",
		},
		{
			"xaabacxcabaaxcabaax", args{"xaabacxcabaaxcabaax"}, "xaabacxcabaax",
		},
		{
			"abbcccbbbcaaccbababcbcabca", args{"abbcccbbbcaaccbababcbcabca"}, "bbcccbb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wants := strings.Split(tt.want, ",")
			got := longestPalindrome(tt.args.s)
			successFlag := false
			for _, want := range wants {
				if got == want {
					successFlag = true
					break

				}
			}
			if !successFlag {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	s10   = ""
	s100  = ""
	s1000 = ""
)

// RandCode 指定指定位数的随机字符串
func RandCode(size int) string {
	code := ""
	chars := "abqw"
	seed := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(seed)
	for i := 0; i < size; i++ {
		index := r2.Intn(len(chars))
		code += string(chars[index])
	}

	return code
}

func init() {
	s10 = RandCode(3) + "abaaba" + RandCode(1)
	s100 = RandCode(50) + "abaabaa" + RandCode(43)
	s1000 = RandCode(500) + "abaabaa" + RandCode(493)
}

func Benchmark_longestPalindrome_10(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome(s10)
	}
}

func Benchmark_longestPalindrome1_10(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome1(s10)
	}
}

func Benchmark_longestPalindrome_100(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome(s100)
	}
}

func Benchmark_longestPalindrome1_100(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome1(s100)
	}
}

func Benchmark_longestPalindrome_1000(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome(s1000)
	}
}

func Benchmark_longestPalindrome1_1000(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome1(s1000)
	}
}
