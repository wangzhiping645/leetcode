package main

func main() {

}

func removeDuplicates(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		nums[index] = nums[i]
		index++
	}
	return index
}
