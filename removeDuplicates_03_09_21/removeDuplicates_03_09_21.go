package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if result != i {
			nums[result] = nums[i]
		}
		result++
	}
	return result
}

func main() {
	fmt.Println(removeDuplicates([]int{2, 2}))
	fmt.Println(removeDuplicates([]int{1, 2, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
