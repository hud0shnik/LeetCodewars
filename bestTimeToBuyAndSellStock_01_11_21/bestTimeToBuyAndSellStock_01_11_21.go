package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minimumPrice, result := prices[0], 0

	for _, price := range prices {
		if price < minimumPrice {
			minimumPrice = price
		} else if (price - minimumPrice) > result {
			result = price - minimumPrice
		}
	}

	return result
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 7, 6, 4, 3, 1, 99}))
}
