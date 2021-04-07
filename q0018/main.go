package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
}

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums)-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j-1 > i && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				if l-1 > j && nums[l] == nums[l-1] {
					l++
					continue
				}
				t := nums[i] + nums[j] + nums[l] + nums[r]
				if t == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				} else if t < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return result
}
