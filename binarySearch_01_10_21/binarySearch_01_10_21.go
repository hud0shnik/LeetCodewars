package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}

	if nums[left] != target {
		return -1
	}

	return left
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(search(items, 63))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, -1))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 12))
}
