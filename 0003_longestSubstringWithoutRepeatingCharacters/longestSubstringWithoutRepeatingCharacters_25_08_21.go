package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	j := 0
	result := 0
	stringMap := make(map[string]bool)

	for i := 0; i < len(s); i++ {
		_, ok := stringMap[string(s[i])]
		if !ok {
			stringMap[string(s[i])] = true
			if (i - j + 1) > result {
				result = i - j + 1
			}
		} else {
			for s[j] != s[i] {
				delete(stringMap, string(s[j]))
				j++
			}
			j++
		}
	}

	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("bcbb"))
}

/*

func lengthOfLongestSubstring(s string) int {
	j := 0
	result := 0
	stringMap := make(map[string]bool)

	for i := 0; i < len(s); i++ {
		_, ok := stringMap[string(s[i])]
		if !ok {
			stringMap[string(s[i])] = true
			result = int(math.Max(float64(result), float64(i-j+1)))
		} else {
			for s[j] != s[i] {
				delete(stringMap, string(s[j]))
				j++
			}
			j++
		}
	}

	return result
}



*/
