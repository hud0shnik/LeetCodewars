package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := 0
	for _, k := range nums {
		if k != 0 {
			nums[i] = k
			i++
		}
	}
	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
}

func main() {
	ex1 := []int{0, 1, 0, 3, 12}
	fmt.Println(ex1)
	moveZeroes(ex1)
	fmt.Println(ex1)
	ex2 := []int{0}
	fmt.Println(ex2)
	moveZeroes(ex2)
	fmt.Println(ex2)
	ex3 := []int{}
	fmt.Println(ex3)
	moveZeroes(ex3)
	fmt.Println(ex3)
}
