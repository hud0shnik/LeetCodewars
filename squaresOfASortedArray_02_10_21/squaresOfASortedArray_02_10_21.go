package main

import "fmt"

func sortedSquares(nums []int) []int {
	numslen := len(nums)
	result := make([]int, numslen)
	k := numslen - 1
	for i, j := 0, k; i <= j; {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[k] = nums[i] * nums[i]
			i++
		} else {
			result[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return result
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
