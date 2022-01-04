package main

import "fmt"

func reverseWords(s string) string {
	result := []rune(s)

	for i := 0; i < len(result); i++ {
		j := i
		for j < len(result) && result[j] != ' ' {
			j++
		}

		for left, right := i, j-1; left < right; left, right = left+1, right-1 {
			result[left], result[right] = result[right], result[left]
		}
		i = j
	}
	return string(result)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("doG gniD"))
}
