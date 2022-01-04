package main

import "fmt"

func lengthOfLastWord(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}
	result := 0
	for i := sLen - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if result != 0 {
				return result
			}
			continue
		}
		result++
	}
	return result
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("a"))
}
