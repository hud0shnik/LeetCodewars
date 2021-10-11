package main

import "fmt"

func combine(n int, k int) [][]int {

	comb := make([]int, k)
	result := [][]int{}

	var recr func(int, int)
	recr = func(x, begin int) {
		if x == k {
			temp := make([]int, k)
			copy(temp, comb)
			result = append(result, temp)
			return
		}

		for i := begin; i <= n+1-k+x; i++ {
			comb[x] = i
			recr(x+1, i+1)
		}
	}

	recr(0, 1)

	return result
}
func main() {
	fmt.Println(combine(4, 2))
	fmt.Println("\t")
	fmt.Println(combine(1, 1))
}
