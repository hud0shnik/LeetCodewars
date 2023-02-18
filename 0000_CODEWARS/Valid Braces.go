package kata

import "strings"

func ValidBraces(str string) bool {
	for strings.Contains(str, "()") || strings.Contains(str, "{}") || strings.Contains(str, "[]") {
		str = strings.ReplaceAll(str, "()", "")
		str = strings.ReplaceAll(str, "{}", "")
		str = strings.ReplaceAll(str, "[]", "")
	}
	return len(str) == 0
}
