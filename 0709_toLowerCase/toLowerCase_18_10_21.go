package main

import "fmt"

func toLowerCase(str string) string {
	result := make([]rune, len(str))
	i := 0
	for _, j := range str {
		if j >= 'A' && j <= 'Z' {
			result[i] = j + rune(32)
		} else {
			result[i] = j
		}

		i++
	}
	return string(result)
}

func main() {
	fmt.Println(toLowerCase("HeLlO wOrLd"))
	fmt.Println(toLowerCase("HeLlO wOrLd!"))
	fmt.Println(toLowerCase("101!"))
}
