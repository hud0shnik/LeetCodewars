package kata

import "strings"

func FirstNonRepeating(str string) string {
	m := map[string]int{}
	for _, r := range str {
		m[strings.ToUpper(string(r))]++
		if m[strings.ToUpper(string(r))] > 1 {
			str = strings.ReplaceAll(str, strings.ToLower(string(r)), "")
			str = strings.ReplaceAll(str, strings.ToUpper(string(r)), "")
		}
	}
	if str == "" {
		return ""
	}
	return string(str[0])
}
