package kata

func Cakes(recipe, available map[string]int) int {
	min := 2147483647
	for i := range recipe {
		if available[i]/recipe[i] < min {
			min = available[i] / recipe[i]
		}
	}

	return min
}
