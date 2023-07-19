package main

import "sort"

func minimumSum(num int) int {

	digits := []int{}

	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	sort.Ints(digits)

	return digits[0]*10 + digits[1]*10 + digits[2] + digits[3]
}
