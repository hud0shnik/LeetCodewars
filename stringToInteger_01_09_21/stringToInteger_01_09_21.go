package main

import "fmt"

func myAtoi(s string) int {
	max := int64(2 << 30)
	signFlag := true
	spaceFlag := true
	sign := 1
	digits := []int{}

	for _, char := range s {
		if char == ' ' && spaceFlag {
			continue
		}
		if signFlag {
			if char == '+' {
				signFlag = false
				spaceFlag = false
				continue
			} else if char == '-' {
				sign = -1
				signFlag = false
				spaceFlag = false
				continue
			}
		}
		if char < '0' || char > '9' {
			break
		}
		spaceFlag, signFlag = false, false
		digits = append(digits, int(char-48))
	}

	var result, place int64
	place, result = 1, 0
	last := -1

	for i, j := range digits {
		if j == 0 {
			last = i
		} else {
			break
		}
	}
	if last > -1 {
		digits = digits[last+1:]
	}

	var rtnMax int64
	if sign > 0 {
		rtnMax = max - 1
	} else {
		rtnMax = max
	}

	digitsLen := len(digits)
	for i := digitsLen - 1; i >= 0; i-- {
		result += int64(digits[i]) * place
		place *= 10
		if digitsLen-i > 10 || result > rtnMax {
			return int(int64(sign) * rtnMax)
		}
	}
	return int(result * int64(sign))
}

func main() {
	fmt.Println(myAtoi("      +3457      "))
	fmt.Println(myAtoi("      -777     dsfcsdfhjl "))
	fmt.Println(myAtoi(" "))
}
