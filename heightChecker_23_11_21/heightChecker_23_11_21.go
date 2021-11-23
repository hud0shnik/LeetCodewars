package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	sortedArray := make([]int, len(heights))
	result := 0
	copy(sortedArray, heights)
	sort.Ints(sortedArray)

	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedArray[i] {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))
	fmt.Println(heightChecker([]int{1, 2, 3, 4, 5}))
}
