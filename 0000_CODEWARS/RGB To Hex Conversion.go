package kata

import "fmt"

func RGB(r, g, b int) string {
	if r > 255 {
		r = 255
	} else if r < 0 {
		r = 0
	}
	if g > 255 {
		g = 255
	} else if g < 0 {
		g = 0
	}
	if b > 255 {
		b = 255
	} else if b < 0 {
		b = 0
	}
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}
