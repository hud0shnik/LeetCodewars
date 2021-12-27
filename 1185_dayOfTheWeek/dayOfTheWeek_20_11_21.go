package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	weekDays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return weekDays[time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Weekday()]
}

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
	fmt.Println(dayOfTheWeek(18, 7, 1999))
	fmt.Println(dayOfTheWeek(15, 8, 1993))
	fmt.Println(dayOfTheWeek(8, 11, 2002))
}
