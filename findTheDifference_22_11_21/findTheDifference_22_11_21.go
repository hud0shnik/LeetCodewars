package main

import "fmt"

func findTheDifference(s string, t string) byte {
	mapa := make(map[rune]int)

	for _, char := range s {
		mapa[char] += 1
	}

	for _, char := range t {
		mapa[char] -= 1
		if mapa[char] < 0 {
			return byte(char)
		}
	}

	return byte(0)
}

func main() {
	fmt.Println(findTheDifference("a", "aa"))
	fmt.Println(findTheDifference("abcd", "abcde"))
	fmt.Println(findTheDifference("", "y"))
	fmt.Println(findTheDifference("ae", "aea"))
}
