package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	xString := strconv.Itoa(x)

	if xString == reverse(xString) {
		return true
	} else {
		return false
	}
}

func reverse(s string) string {
	reversed := []rune(s)
	for i, j := 0, len(reversed)-1; i < len(reversed)/2; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}

func main() {
	fmt.Println(isPalindrome(-101))
}

///////*-*-*---*---*-*-*-*-*-*----*-*-**----*-*---*-*-*
/*

func isPalindrome(x int) bool {
    if x>=0{
        return x == reverse(x)
    }
    return false
}

func reverse(x int) int {
    rev := 0
    for x!= 0{
        num := x%10
        rev = rev * 10 + num
        x = x/10
    }
    return rev
}

*/
