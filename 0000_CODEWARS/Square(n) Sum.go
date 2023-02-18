package kata

func SquareSum(numbers []int) int {
	res := 0
	for _, n := range numbers {
		res += n * n
	}
	return res
}
