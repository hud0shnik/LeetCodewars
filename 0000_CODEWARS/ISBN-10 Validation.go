package kata

import (
	"strconv"
	"strings"
)

func ValidISBN10(isbn string) bool {
	digits := strings.Split(isbn, "")
	sum := 0
	if len(digits) != 10 {
		return false
	}
	for i := 0; i < 10; i++ {
		a, err := strconv.Atoi(digits[i])
		if err != nil {
			if i == 9 && (digits[i] == "X" || digits[i] == "x") {
				sum += 100
			} else {
				return false
			}
		} else {
			sum += (i + 1) * a
		}
	}

	return sum%11 == 0
}
