package kata

import "strings"

func duplicate_count(s1 string) int {
	s := strings.ToLower(s1)
	m := map[rune]int{}
	res := 0
	for _, r := range s {
		m[r]++
		if m[r] == 2 {
			res++
		}
	}
	return res
}
