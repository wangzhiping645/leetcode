package main

import "fmt"

func main() {

	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-1534236469))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1563847412))

}

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + (x % 10)
		x = x / 10
	}
	if result < -1<<31 || result > 1<<31-1 {
		result = 0
	}
	return result
}
