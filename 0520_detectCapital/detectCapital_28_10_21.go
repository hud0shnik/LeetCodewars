package main

import (
	"fmt"
	"strings"
)

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	if word[1:] == strings.ToLower(word[1:]) {
		return true
	}
	if word == strings.ToUpper(word) {
		return true
	}
	return false
}

func main() {
	fmt.Println(detectCapitalUse("AAA"))
	fmt.Println(detectCapitalUse("UwU"))
	fmt.Println(detectCapitalUse("Owo"))
	fmt.Println(detectCapitalUse("uwu"))
}
