package kata

import (
	"strconv"
	"strings"
	"unicode"
)

func Order(sentence string) string {

	words := strings.Fields(sentence)
	res := make([]string, len(words))
	for _, w := range words {
		for _, r := range w {
			if unicode.IsDigit(r) {
				i, _ := strconv.Atoi(string(r))
				res[i-1] = w
			}
		}
	}

	return strings.Join(res, " ")
}
