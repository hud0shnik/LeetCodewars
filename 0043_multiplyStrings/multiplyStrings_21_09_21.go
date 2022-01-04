package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" {
		return "0"
	}
	if num2 == "0" {
		return "0"
	}

	rune1 := []rune(num1)
	rune2 := []rune(num2)
	t := make([]int, len(rune1)+len(rune2))

	for i := 0; i < len(rune1); i++ {
		for j := 0; j < len(rune2); j++ {
			t[i+j+1] += int(rune1[i]-'0') * int(rune2[j]-'0')
		}
	}

	for i := len(t) - 1; i > 0; i-- {
		t[i-1] += t[i] / 10
		t[i] = t[i] % 10
	}
	if t[0] == 0 {
		t = t[1:]
	}

	result := make([]rune, len(t))
	for i := 0; i < len(t); i++ {
		result[i] = '0' + rune(t[i])
	}

	return string(result)
}

func main() {
	fmt.Println(multiply("3", "4"))
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("123", "0"))
	fmt.Println(multiply("199999999999999999999999", "99999999999999999999999999999999999999993245"))
}
