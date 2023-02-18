package kata

// "(())((()())())()"
// '())('
// '((()()))'
func ValidParentheses(parens string) bool {
	if len(parens) == 0 {
		return true
	}
	if parens[0] == ')' || len(parens)%2 != 0 {
		return false
	}
	n := 0

	for _, p := range parens {
		if p == '(' {
			n++
		}
		if p == ')' {
			n--
			if n <= -1 {
				return false
			}
		}
	}

	return n == 0
}
