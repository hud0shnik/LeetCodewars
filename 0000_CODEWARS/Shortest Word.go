package kata

import "strings"

func FindShort(s string) int {
	sl := strings.Fields(s)
	min := len(sl[0])
	for _, w := range sl {
		if len(w) < min {
			min = len(w)
		}
	}
	return min
}
