package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	matRow := len(mat[0])
	size := len(mat) * matRow

	if size != r*c {
		return mat
	}

	var result [][]int

	for j, i := 0, 0; j < r && i < size; j++ {
		resultRow := make([]int, c)
		for k := 0; k < c && i < size; k++ {
			resultRow[k] = mat[i/matRow][i%matRow]
			i++
		}
		result = append(result, resultRow)
	}

	return result
}

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
}
