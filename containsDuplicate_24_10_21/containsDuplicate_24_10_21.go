package main

import "fmt"

func containsDuplicate(nums []int) bool {
	mapa := map[int]bool{}

	for _, num := range nums {
		if mapa[num] {
			return true
		}

		mapa[num] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 4, 5, 3, 2}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
