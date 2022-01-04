package main

import "fmt"

func permute(nums []int) [][]int {
	lenNums := len(nums)
	path := make([]int, lenNums)
	used := make([]bool, lenNums)
	result := make([][]int, 0)
	var recr func(int)

	recr = func(u int) {
		if u == lenNums {
			t := make([]int, lenNums)
			copy(t, path)
			result = append(result, t)
			return
		}
		for i := 0; i < lenNums; i++ {
			if !used[i] {
				path[u] = nums[i]
				used[i] = true
				recr(u + 1)
				used[i] = false
			}
		}
	}

	recr(0)
	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
