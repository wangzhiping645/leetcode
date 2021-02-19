package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
}

// 最优解，双指针 + 排序
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	current := abs(nums[0] + nums[1] + nums[2] - target)

	for i := 0; i < len(nums)-2; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if l-1 > i && nums[l] == nums[l-1] {
				l++
				continue
			}
			t := nums[i] + nums[l] + nums[r] - target
			if t == 0 {
				return target
			}
			if current > abs(t) {
				result = nums[i] + nums[l] + nums[r]
				current = abs(t)
			}
			if t < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
