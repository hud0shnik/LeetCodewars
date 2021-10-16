package main

import "fmt"

func numSquares(n int) int {

	result := make([]int, n+1)
	result[0] = 0

	for i := 1; i <= n; i++ {
		result[i] = i

		for j := 1; j*j <= i; j++ {
			result[i] = min(result[i], result[i-j*j]+1)
		}
	}

	return result[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(14))
	fmt.Println(numSquares(3))
}
