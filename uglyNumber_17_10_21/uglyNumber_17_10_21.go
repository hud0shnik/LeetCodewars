package main

import "fmt"

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num <= 6 {
		return true
	}

	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(14))
	fmt.Println(isUgly(1024))
}
