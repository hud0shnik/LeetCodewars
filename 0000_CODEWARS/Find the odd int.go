package kata

func FindOdd(seq []int) int {
	m := map[int]int{}
	for _, n := range seq {
		m[n]++
	}
	for i, j := range m {
		if j%2 != 0 {
			return i
		}
	}
	return -1
}
