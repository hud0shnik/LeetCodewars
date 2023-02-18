package kata

func FindUniq(arr []float32) float32 {
	m := map[float32]int{}
	for _, f := range arr {
		m[f]++
	}
	for f, n := range m {
		if n == 1 {
			return f
		}
	}
	return -1
}
