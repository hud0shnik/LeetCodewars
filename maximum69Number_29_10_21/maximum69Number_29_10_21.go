package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	result, _ := strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return result
}

func main() {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9999))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(6666))
}
