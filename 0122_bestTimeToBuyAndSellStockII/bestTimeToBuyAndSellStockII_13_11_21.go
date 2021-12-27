package main

import "fmt"

func maxProfit(prices []int) int {
	result := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			result += prices[i+1] - prices[i]
		}
	}

	return result
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
