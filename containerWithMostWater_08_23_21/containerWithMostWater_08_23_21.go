package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0

	for left < right {
		currentMax := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		result = int(math.Max(float64(result), float64(currentMax)))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 2, 1}))
}
