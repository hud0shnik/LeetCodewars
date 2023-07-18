package main

func numIdenticalPairs(nums []int) int {

	res := 0
	m := map[int]int{}

	for _, n := range nums {
		res += m[n]
		m[n]++
	}

	return res
}
