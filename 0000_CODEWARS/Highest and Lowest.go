package kata

import (
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	nums := strings.Split(in, " ")
	var min, max int = 2147483647, -2147483647
	for _, a := range nums {
		i, _ := strconv.Atoi(a)
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	return strconv.Itoa(max) + " " + strconv.Itoa(min)
}
