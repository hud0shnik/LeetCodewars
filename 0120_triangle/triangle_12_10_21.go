package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	len1 := len(triangle)
	result := make([]int, len1+1)
	for i := len1 - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			result[j] = min(result[j], result[j+1]) + triangle[i][j]
		}
	}
	return result[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3}}))
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 1},
		{6, 5, 1},
		{4, 1, 8, 3}}))
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 1},
		{1, 5, 3},
		{1, 1, 8, 3}}))
}
