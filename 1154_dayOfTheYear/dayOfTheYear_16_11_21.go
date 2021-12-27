package main

import (
	"fmt"
	"time"
)

func dayOfYear(date string) int {
	dateTime, _ := time.Parse("2006-01-02", date)

	return int(dateTime.Sub(time.Date(dateTime.Year(), 1, 1, 0, 0, 0, 0, time.UTC)).Round(time.Hour).Hours())/24 + 1
}

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
	fmt.Println(dayOfYear("2019-02-10"))
	fmt.Println(dayOfYear("2003-03-01"))
	fmt.Println(dayOfYear("2004-03-01"))

	fmt.Println("\t today: ", dayOfYear(time.Now().Format("2006-01-02")))
}
