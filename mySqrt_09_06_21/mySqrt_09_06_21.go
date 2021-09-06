package main

import "fmt"

func mySqrt(x int) int {
	left := 0
	right := x
	for left < right {
		middle := (left + right + 1) / 2
		if middle*middle > x {
			right = middle - 1
		} else {
			left = middle
		}
	}
	return left
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
}
