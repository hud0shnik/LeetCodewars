package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	var result [][]int

	for i := 0; i < numRows; i++ {
		resultRow := make([]int, i+1)
		resultRow[0] = 1
		resultRow[i] = 1
		if i > 1 {
			for j := 1; j < i; j++ {
				resultRow[j] = result[i-1][j-1] + result[i-1][j]
			}
		}

		result = append(result, resultRow)
	}

	return result
}

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}
