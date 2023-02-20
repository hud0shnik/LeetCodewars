package kata

import (
	"unicode"
)

func alphanumeric(str string) bool {
	for _, p := range str {
		if !(unicode.IsDigit(p) || unicode.Is(unicode.Latin, p)) {
			return false
		}
	}
	return str != ""
}
