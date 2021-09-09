package main

import "fmt"

func addBinary(a string, b string) string {
	var carry, sum rune = '0', '0'
	aRune, bRune := []rune(a), []rune(b)
	small, big := aRune, bRune

	if len(aRune) > len(bRune) {
		small, big = bRune, aRune
	}

	result := make([]rune, len(big)+1)
	newSmall := make([]rune, len(big)-len(small))

	for i := range newSmall {
		newSmall[i] = '0'
	}

	newSmall = append(newSmall, small...)
	for i := len(big) - 1; i >= 0; i-- {
		sum, carry = add(newSmall[i], big[i], carry)
		result[i+1] = sum
	}

	if carry == '1' {
		result[0] = carry
	} else {
		result = result[1:]
	}
	return string(result)
}

func add(a, b, c rune) (rune, rune) {
	if a == '0' && b == '0' {
		return c, '0'
	}
	if a == '1' && b == '1' {
		return c, '1'
	}
	if c == '1' {
		return '0', '1'
	}
	return '1', '0'
}

func main() {
	fmt.Println(addBinary("101", "1"))
}
