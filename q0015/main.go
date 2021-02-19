package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}

// threeSum1 解法一
// 排序 + 遍历后获取各个元素的数量map
// 双层遍历 + 计算与target的差值，如果该差值在map中存在 且次数足够，则符合要求
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j-1 > i && nums[j] == nums[j-1] {
				continue
			}
			need := 0 - nums[i] - nums[j]
			v, found := m[need]
			if !found {
				continue
			}
			if need < nums[i] || need < nums[j] {
				continue
			}
			times := 0
			if need == nums[i] {
				times++
			}
			if need == nums[j] {
				times++
			}
			if v > times {
				result = append(result, []int{nums[i], nums[j], need})
			}
		}
	}

	return result
}

// 最优解，双指针 + 排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

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
			t := nums[i] + nums[l] + nums[r]
			if t == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if t < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return result
}
