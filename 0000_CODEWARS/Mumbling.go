package kata

import "strings"

func Accum(s string) string {
	res := ""
	j := 0
	for _, ch := range s {
		res += strings.ToUpper(string(ch))
		for i := 0; i < j; i++ {
			res += strings.ToLower(string(ch))
		}
		res += "-"
		j++
	}
	return res[:len(res)-1]
}
