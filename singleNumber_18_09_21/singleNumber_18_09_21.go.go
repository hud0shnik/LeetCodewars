package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i] // ^ - XOR
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2, 4, 0}))
	fmt.Println(singleNumber([]int{1}))
}
