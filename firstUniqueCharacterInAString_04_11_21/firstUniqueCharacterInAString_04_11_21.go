package main

import "fmt"

func firstUniqChar(s string) int {
	runee := [27]rune{}
	len := len(s)

	for i := 0; i < len; i++ {
		runee[s[i]-'a']++
	}
	for i := 0; i < len; i++ {
		if runee[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
