package main

import "fmt"

func majorityElement(nums []int) int {
	mapa := make(map[int]int)
	for _, i := range nums {
		mapa[i]++
		if mapa[i] > len(nums)/2 {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{}))
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{3, 2, 3, 2})) // output: 0 :^/
}
