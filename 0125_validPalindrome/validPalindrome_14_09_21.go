package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !(('a' <= s[i] && s[i] <= 'z') || ('0' <= s[i] && s[i] <= '9')) {
			i++
		}
		for i < j && !(('a' <= s[j] && s[j] <= 'z') || ('0' <= s[j] && s[j] <= '9')) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("rtrtt"))
	fmt.Println(isPalindrome("t t y t t"))
	fmt.Println(isPalindrome("0P"))
}
