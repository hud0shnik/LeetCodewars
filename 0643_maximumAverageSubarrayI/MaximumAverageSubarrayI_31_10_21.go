package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < 2 {
		return float64(nums[0])
	}
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maximum := sum
	pointer := k

	for pointer < len(nums) {
		sum = sum - nums[pointer-k] + nums[pointer]
		if sum > maximum {
			maximum = sum
		}
		pointer++
	}

	return float64(maximum) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}
