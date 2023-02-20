package kata

import (
	"fmt"
	"strings"
)

func FormatDuration(seconds int64) string {
	args := []string{}
	m, s := divmode(seconds, 60)
	h, m := divmode(m, 60)
	d, h := divmode(h, 24)
	y, d := divmode(d, 365)

	time := []int64{y, d, h, m, s}
	fmt.Println(time)
	words := []string{"year", "day", "hour", "minute", "second"}

	for i, t := range time {
		if t == 1 {
			args = append(args, fmt.Sprintf("1 %s", words[i]))
		} else if t > 1 {
			args = append(args, fmt.Sprintf("%d %ss", t, words[i]))
		}
	}

	if len(args) == 1 {
		return args[0]
	}
	if len(args) == 2 {
		return fmt.Sprintf("%s and %s", args[0], args[1])
	}
	return strings.Join(args[:len(args)-1], ", ") + " and " + args[len(args)-1]

}

func divmode(a, b int64) (int64, int64) {
	return a / b, a % b
}
