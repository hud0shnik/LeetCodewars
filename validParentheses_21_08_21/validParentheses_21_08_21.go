package main

import "fmt"

func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	stack := make([]string, len(s))
	id := 0

	for i := 0; i < len(s); i++ {
		char := string(s[i])

		if char == "(" || char == "{" || char == "[" {
			stack[id] = char
			id++
		} else {

			if id == 0 {
				return false
			}

			id--
			char2 := stack[id]

			if (char == ")" && char2 != "(") || (char == "}" && char2 != "{") || (char == "]" && char2 != "[") {
				return false
			}
		}
	}

	return id == 0
}

func main() {
	fmt.Println(isValid("({(}))"))
	fmt.Println(isValid("({()})"))
}
