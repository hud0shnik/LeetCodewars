package main

import "fmt"

func fib(N int) int {
	if N <= 1 {
		return N
	}

	a, b := 0, 1
	for i := 2; i <= N; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(19))
}
