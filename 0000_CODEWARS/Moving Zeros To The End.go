package kata

func MoveZeros(arr []int) []int {
	res := []int{}
	zeros := 0
	for _, num := range arr {
		if num == 0 {
			zeros++
		} else {
			res = append(res, num)
		}
	}
	for i := 0; i < zeros; i++ {
		res = append(res, 0)
	}
	return res
}
