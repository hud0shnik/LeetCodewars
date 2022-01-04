package main

import "fmt"

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	for i := 0; i < haystackLen-needleLen+1; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
