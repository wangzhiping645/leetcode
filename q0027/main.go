package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	x := removeElement1(a, 2)
	fmt.Println(a, x)
}

func removeElement(nums []int, val int) int {
	result := 0
	for _, value := range nums {
		if value != val {
			nums[result] = value
			result++
		}
	}
	return result
}

// 内存最优
func removeElement1(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != result {
				nums[result], nums[i] = nums[i], nums[result]
			}
			result++
		}
	}
	return result
}
