package main

import "fmt"

func main() {
	//fmt.Print(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
	fmt.Print(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	//fmt.Print(maxArea([]int{1, 1}))

}

func maxArea1(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	for i < j {
		t := min(height[i], height[j]) * (j - i)
		if t > max {
			max = t
		}
		if height[j] < height[i] {
			j--
		} else {
			i++
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	max := 0
	h := 0
	i, j := 0, len(height)-1
	for i < j {
		if height[i] < h {
			i++
			continue
		}
		if height[j] < h {
			j--
			continue
		}
		if height[j] < height[i] {
			if max < height[j]*(j-i) {
				max = height[j] * (j - i)
				h = height[j]
			}
			j--
		} else {
			if max < height[i]*(j-i) {
				max = height[i] * (j - i)
				h = height[i]
			}
			i++
		}
	}
	return max
}
