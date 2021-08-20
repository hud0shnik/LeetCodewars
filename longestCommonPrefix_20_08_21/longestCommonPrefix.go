package main

import (
	"fmt"
	"time"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	last := 0
	minLen := len(strs[0])
	result := ""

	for i, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
			last = i
		}
	}

	for i := 0; i < minLen; i++ {
		match := true
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != strs[last][i] {
				match = false
			}
		}
		if match {
			result += string(strs[last][i])
		} else {
			break
		}
	}
	return result
}

func main() {
	stressStr := []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaq", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaq", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaq", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaq", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz"}
	/*str := []string{"flower", "flow", "flight", "flynn"}
	fmt.Println(longestCommonPrefix(str))
	fmt.Println(longestCommonPrefix([]string{"arcds", "ard", ""}))
	fmt.Println(longestCommonPrefix([]string{"a", "asdfdsf"}))*/

	fmt.Println("Начать стресстест!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	start := time.Now()
	for k := 0; k < 100; k++ {
		fmt.Println(longestCommonPrefix(stressStr))
	}
	duration := time.Since(start)
	fmt.Println("\ttime:", duration)
}
