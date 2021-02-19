package main

import (
	"math/rand"
	"time"
)

var (
	str = ""
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
	str = RandCode(10000)
}
