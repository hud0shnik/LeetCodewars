package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	priceLen := len(prices)
	result := make([]int, priceLen+1)
	minimum := prices[0]

	for i := 1; i < priceLen+1; i++ {
		if i >= 2 {
			minimum = Min(minimum, prices[i-1]-result[i-2])
		}
		result[i] = Max(result[i-1], prices[i-1]-minimum)
	}

	return result[priceLen]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}
