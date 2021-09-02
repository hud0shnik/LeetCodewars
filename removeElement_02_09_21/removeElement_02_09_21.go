package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != result {
				nums[i], nums[result] = nums[result], nums[i]
			}
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(removeElement([]int{1, 2}, 2))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
