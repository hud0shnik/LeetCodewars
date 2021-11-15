package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	lenS := len(s)
	results := make([]bool, lenS+1)
	results[0] = true

	for _, w := range wordDict {
		words[w] = true
	}

	for i := 1; i <= lenS; i++ {
		for j := 0; j < i; j++ {
			if results[j] && words[s[j:i]] {
				results[i] = true
				break
			}
		}
	}

	return results[lenS]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
