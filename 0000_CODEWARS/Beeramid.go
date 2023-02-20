package kata

func Beeramid(bonus int, price float64) int {
	var res int
	for i, j := 0, 1; ; j++ {
		i += j * j
		if float64(i)*price > float64(bonus) {
			break
		}
		res++
	}
	return res
}
