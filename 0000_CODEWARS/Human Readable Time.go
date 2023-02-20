package kata

import "fmt"

func HumanReadableTime(sec int) string {
	return fmt.Sprintf("%.2d:%.2d:%.2d", sec/3600, sec%3600/60, sec%60)
}
