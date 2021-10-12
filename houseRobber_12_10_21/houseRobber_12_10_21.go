package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	oldResult, result := 0, 0
	for i := 0; i < len(nums); i++ {
		oldResult, result = result, max(result, oldResult+nums[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))

	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1, 3, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1, 31, 1, 1, 37, 1}))
}
