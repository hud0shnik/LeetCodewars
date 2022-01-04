package main

import "fmt"

func recursionSearch(nums []int, left, right, target int) int {
	if left >= right {
		if nums[left] == target {
			return left
		}
		return -1
	}
	if nums[left] < nums[right] {
		if nums[left] > target || nums[right] < target {
			return -1
		}
		middle := int(left + (right-left)/2)
		rec := recursionSearch(nums, left, middle, target)
		if rec != -1 {
			return rec
		}
		return recursionSearch(nums, middle+1, right, target)
	} else {
		middle := int(left + (right-left)/2)
		rec := recursionSearch(nums, left, middle, target)
		if rec != -1 {
			return rec
		}
		return recursionSearch(nums, middle+1, right, target)
	}
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	return !(-1 == recursionSearch(nums, 0, len(nums)-1, target))
}

func main() {
	fmt.Println(search([]int{8, 9, 10, 11, 1, 2, 3, 4, 5, 6, 7}, 11))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 0, 1, 1}, 0))
}
