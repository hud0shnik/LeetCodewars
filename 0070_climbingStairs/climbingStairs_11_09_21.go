package main

import "fmt"

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	result := make([]int, n+1)
	result[0] = 1
	result[1] = 1
	for i := 2; i < n+1; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}

func main() {
	for i := 0; i <= 32; i++ {
		fmt.Println(i, ":\t", climbStairs(i))
	}
}
