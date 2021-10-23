package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		middle := left + (right-left)/2

		if nums[left] > nums[middle] {
			right = middle
		} else if nums[middle] > nums[right] {
			left = middle + 1
		} else if nums[middle] == nums[right] {
			right--
		} else {
			right = left
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{1, 3, 5}))
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

func findMin1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if result > nums[i] {
			return nums[i]
		}
	}
	return result
}
