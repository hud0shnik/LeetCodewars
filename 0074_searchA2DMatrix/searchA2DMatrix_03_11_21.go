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

func searchMatrix(matrix [][]int, target int) bool {
	size := len(matrix[0]) * len(matrix)
	reshapedMatrix := matrixReshape(matrix, 1, size)
	for i := 0; i < size; i++ {
		if reshapedMatrix[0][i] == target {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 60))
}
