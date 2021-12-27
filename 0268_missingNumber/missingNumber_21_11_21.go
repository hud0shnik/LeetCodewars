package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{0}))
}
