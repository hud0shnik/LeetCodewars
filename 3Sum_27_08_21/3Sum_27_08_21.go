package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			resultSum := nums[i] + nums[j] + nums[k]
			if resultSum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if resultSum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return result
}

func main() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))

}
