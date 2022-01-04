package main

import "fmt"

func letterCasePermutation(s string) []string {
	result := []string{}
	var recr func([]byte, string, int)

	recr = func(str []byte, s string, u int) {
		if len(str) == len(s) {
			result = append(result, string(str))
			return
		}

		if u < len(s) {
			recr(append(str, s[u]), s, u+1)
			if s[u] >= 'a' && s[u] <= 'z' {
				recr(append(str, s[u]-32), s, u+1)
			} else if s[u] >= 'A' && s[u] <= 'Z' {
				recr(append(str, s[u]+32), s, u+1)
			}

		}
	}

	recr([]byte{}, s, 0)
	return result
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
