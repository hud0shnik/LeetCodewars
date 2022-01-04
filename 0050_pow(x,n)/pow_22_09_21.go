package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n <= 0 {
		return 1.0 / powRec(x, -n)
	}
	return powRec(x, n)
}

func powRec(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	result := powRec(x, n>>1) // Left shift

	if n&1 == 0 { // Bitwise AND
		return result * result
	}

	return result * result * x
}

func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(0, 10))
	fmt.Println(myPow(0, 0))
	fmt.Println(myPow(1024, -10))
}
