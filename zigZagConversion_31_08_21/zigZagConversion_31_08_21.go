package main

import "fmt"

func convert(s string, numRows int) string {
	down := 0
	up := numRows - 2
	matrix := make([][]rune, numRows, numRows)

	for i := 0; i != len(s); i++ {
		if down != numRows {
			matrix[down] = append(matrix[down], rune(s[i]))
			down++
		} else if up > 0 {
			matrix[up] = append(matrix[up], rune(s[i]))
			up--
		} else {
			up = numRows - 2
			down = 0
			i--
		}
	}
	result := make([]rune, 0, len(s))

	for _, row := range matrix {
		for _, item := range row {
			result = append(result, item)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("A.", 1))
}
