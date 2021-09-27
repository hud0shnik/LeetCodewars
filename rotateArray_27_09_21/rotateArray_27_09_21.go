package main

import "fmt"

func rotate(nums []int, k int) {
	numsLen := len(nums)
	result := make([]int, numsLen)

	for i := range nums {
		result[(i+k)%numsLen] = nums[i]
	}

	copy(nums, result)

}

func main() {
	ex := []int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		rotate(ex, i)
		fmt.Println(ex)
	}

}
