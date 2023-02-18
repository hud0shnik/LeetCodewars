package kata

import (
	"sort"
	"strconv"
	"strings"
)

func OrderWeight(strng string) string {
	list := strings.Fields(strng)
	sort.Slice(list, func(a, b int) bool {
		sumA, sumB := 0, 0
		for _, r := range list[a] {
			i, _ := strconv.Atoi(string(r))
			sumA += i
		}
		for _, r := range list[b] {
			i, _ := strconv.Atoi(string(r))
			sumB += i
		}
		if sumA < sumB {
			return true
		}
		if sumA == sumB {
			return list[a] < list[b]
		}
		return false

	})

	return strings.Join(list, " ")
}

/*


func OrderWeight(strng string) string {
	list := strings.Fields(strng)
	m := map[int]string{}
	res := []string{}
	for _, l := range list {
		sum := 0
		for _, num := range l {
			i, _ := strconv.Atoi(string(num))
			sum += i
		}
		m[sum] = l
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		res = append(res, m[k])
	}

	return strings.Join(res, " ")
}
*/
