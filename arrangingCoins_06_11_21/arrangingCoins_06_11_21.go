package main

import "fmt"

func arrangeCoins(n int) int {
	result, coins, i := 0, n, 0

	for ; ; result++ {
		i++
		if !(coins-i >= 0) {
			break
		}
		coins -= i
	}

	return result
}

func main() {
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(7775))
}
