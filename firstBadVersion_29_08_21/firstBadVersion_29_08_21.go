package main

import "fmt"

func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		middle := left + (right-left)/2
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	m := make(map[int]bool)
	m[3] = false
	m[4] = true
	m[5] = false

	if v, ok := m[version]; ok {
		return v
	}

	return false
}

func main() {
	fmt.Println(5, 4)
	fmt.Println(1, 1)
}
