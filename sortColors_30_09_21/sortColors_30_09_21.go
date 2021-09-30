package main

import "fmt"

func sortColors(nums []int) {
	left, current := 0, 0
	right := len(nums) - 1

	for current <= right {
		if nums[current] == 0 {
			nums[left], nums[current] = nums[current], nums[left]
			left++
			current++
		} else if nums[current] == 2 {
			nums[current], nums[right] = nums[right], nums[current]
			right--
		} else {
			current++
		}
	}
}

func main() {
	n1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(n1)
	fmt.Println(n1)

	n2 := []int{2, 0, 1}
	sortColors(n2)
	fmt.Println(n2)

	n3 := []int{0}
	sortColors(n3)
	fmt.Println(n3)

	n4 := []int{1}
	sortColors(n4)
	fmt.Println(n4)
}
