package kata

func GetMiddle(s string) string {
	l := len(s)
	if l%2 == 0 {
		return string(s[l/2-1]) + string(s[l/2])
	}
	return string(s[(l-1)/2])
}
