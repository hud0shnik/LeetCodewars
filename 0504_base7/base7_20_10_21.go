package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	result, counter := 0, 1

	for num > 0 || num < 0 {
		result += (num % 7) * counter
		num /= 7
		counter *= 10
	}

	return strconv.Itoa(result)
}

func main() {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
}
