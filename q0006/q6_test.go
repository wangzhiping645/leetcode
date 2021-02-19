package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_convert2(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"PAYPALISHIRING", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert2(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	s = ""
)

// RandCode 指定指定位数的随机字符串
func RandCode(size int) string {
	code := ""
	chars := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	seed := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(seed)
	for i := 0; i < size; i++ {
		index := r2.Intn(len(chars))
		code += string(chars[index])
	}
	return code
}

func init() {
	s = RandCode(1000)
}

const (
	column = 50
)

func Benchmark_convert(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convert(s, column)
	}
}

func Benchmark_convert1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convert1(s, column)
	}
}

func Benchmark_convert2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convert2(s, column)
	}
}
