package main

import "fmt"

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	word1 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(word1)
	reverseString(word1)
	fmt.Println(word1)
}
