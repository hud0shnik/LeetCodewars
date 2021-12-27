package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	return dividend / divisor
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(-2147483648, -1))
	fmt.Println(divide(8888, 8888888))
}
