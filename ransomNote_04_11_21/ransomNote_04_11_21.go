package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	len1 := len(ransomNote)
	len2 := len(magazine)

	if len1 == 0 && len2 == 0 {
		return true
	}
	if len1 > len2 {
		return false
	}

	letters := make([]int, 26)

	for _, i := range ransomNote {
		letters[i-'a']++
	}
	for _, i := range magazine {
		letters[i-'a']--
	}

	for _, i := range letters {
		if i > 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
