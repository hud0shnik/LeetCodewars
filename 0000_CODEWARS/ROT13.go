package kata

func Rot13(msg string) string {
	var res []rune
	for _, r := range msg {
		if r >= 'a' && r <= 'z' {
			if r > 'm' {
				res = append(res, r-13)
			} else {
				res = append(res, r+13)
			}
		} else if r >= 'A' && r <= 'Z' {
			if r > 'M' {
				res = append(res, r-13)
			} else {
				res = append(res, r+13)
			}
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}
