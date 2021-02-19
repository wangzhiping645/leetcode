package main

import "fmt"

func main() {
	//fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{}))
	//fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{}))
	//fmt.Println(findMedianSortedArrays([]int{}, []int{1, 2}))
	//fmt.Println(findMedianSortedArrays([]int{}, []int{1, 2, 3}))
	//fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5}))
	//fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5}))
	//fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{1, 2, 3}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	fmt.Println(nums1Mid, nums2Mid, k)
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 && l2 == 0 {
		return 0
	}
	if l2 == 0 {
		if l1%2 == 0 {
			return (float64(nums1[l1/2-1]) + float64(nums1[l1/2])) / 2
		} else {
			return float64(nums1[l1/2])
		}
	}

	if l1 == 0 {
		if l2%2 == 0 {
			return (float64(nums2[l2/2-1]) + float64(nums2[l2/2])) / 2
		} else {
			return float64(nums2[l2/2])
		}
	}

	result := 0
	if (l1+l2)%2 == 1 {
		target := (l1+l2)/2 + 1
		current := 0
		i, j := 0, 0
		for current < target {
			//fmt.Println(i, j, current)
			current++
			if i >= l1 {
				result = nums2[j]
				j++
				continue
			}
			if j >= l2 {
				result = nums1[i]
				i++
				continue
			}
			if nums1[i] <= nums2[j] {
				if i < l1 {
					result = nums1[i]
					i++
				} else {
					result = nums2[j]
					j++
				}
			} else {
				if j < l2 {
					result = nums2[j]
					j++
				} else {
					result = nums1[i]
					i++
				}
			}
		}
		return float64(result)
	} else {
		//target1 := (l1+l2)/2
		target2 := (l1+l2)/2 + 1
		result1 := 0
		result2 := 0
		current := 0
		i, j := 0, 0
		for current < target2 {
			//fmt.Println(i, j, current)
			result1 = result2
			current++
			if i >= l1 {
				result2 = nums2[j]
				j++
				continue
			}
			if j >= l2 {
				result2 = nums1[i]
				i++
				continue
			}
			if nums1[i] <= nums2[j] {
				if i < l1 {
					result2 = nums1[i]
					i++
				} else {
					result2 = nums2[j]
					j++
				}
			} else {
				if j < l2 {
					result2 = nums2[j]
					j++
				} else {
					result2 = nums1[i]
					i++
				}
			}
		}

		return (float64(result1) + float64(result2)) / 2
	}

}
