package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {

	result := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect([]int{4, 9, 5, 3, 2, 1}, []int{9, 4, 9, 8, 4, 5, 3, 2, 1}))
}
