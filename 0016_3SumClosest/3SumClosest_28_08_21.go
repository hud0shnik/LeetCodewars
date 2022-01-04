package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	result := 0
	closest := 10000
	absSum := 0

	sort.Ints(nums)

	for k := 0; k < len(nums)-2; k++ {
		i := k + 1
		j := len(nums) - 1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if target-sum < 0 {
				absSum = -(target - sum)
			} else {
				absSum = target - sum
			}
			if absSum < closest {
				closest = absSum
				result = sum
			}
			if nums[k]+nums[i]+nums[j] > target {
				j--
			} else {
				i++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{-1, 23, 10, -5}, -4))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
