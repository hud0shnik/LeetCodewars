package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	for num >= 10 {
		num = num/10 + num%10
	}

	return num
}

func main() {
	fmt.Println(addDigits(38))
	fmt.Println(addDigits(0))
}
