package kata

import "strconv"

func Persistence(n int) int {

	res := 0
	for ; n >= 10; res++ {
		str := strconv.Itoa(n)
		n = 1
		for _, c := range str {
			i, _ := strconv.Atoi(string(c))
			n *= i
		}
	}
	return res

}
