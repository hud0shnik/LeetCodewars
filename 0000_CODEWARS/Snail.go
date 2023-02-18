package kata

// 1 2 3
// 4 5 6
// 7 8 9

func Snail(matrix [][]int) []int {

	size := len(matrix)

	if len(matrix[0]) == 0 {
		return []int{}
	}

	var res = make([]int, size*size)
	resJ := -1
	for i := 0; i <= size-i; i++ {
		for j := i; j <= size-i-1; j++ {
			resJ++
			res[resJ] = matrix[i][j]
		}

		if i == size-i-1 {
			break
		}

		for j := i + 1; j <= size-i-2; j++ {
			resJ++
			res[resJ] = matrix[j][size-i-1]
		}

		for j := size - i - 1; j >= i; j-- {
			resJ++
			res[resJ] = matrix[size-i-1][j]
		}

		for j := size - i - 2; j >= i+1; j-- {
			resJ++
			res[resJ] = matrix[j][i]
		}
	}

	return res
}
