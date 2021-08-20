package main

import "fmt"

func reverse(x int) int {
	if x < 0 {
		newx := -x
		newint := 0
		for newx > 0 {
			remainder := newx % 10
			newint *= 10
			newint += remainder
			newx /= 10
		}
		if newint > 2147483647 {
			return 0
		}
		return -newint
	}
	newint := 0
	for x > 0 {
		remainder := x % 10
		newint *= 10
		newint += remainder
		x /= 10
	}
	if newint > 2147483647 {
		return 0
	}
	return newint
}

func main() {
	fmt.Println(reverse(-1010))
}

/*
func reverse0(x int) int {
	if x > 0 {
		xString := strconv.Itoa(x)
		runes := []rune(xString)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result, err := strconv.ParseInt(string(runes), 0, 64)
		if err == nil {
			return int(result)
		} else {
			return 0
		}
	} else if x < 0 {
		xString := strconv.Itoa(-x)
		runes := []rune(xString)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result, err := strconv.ParseInt(string(runes), 0, 64)
		if err == nil {
			return int(-result)
		} else {
			return 0
		}

	} else {
		return 0
	}
}*/

/*
func reverse(x int) int {
       if -10 < x && x < 10 {
		return x
	}

	isNegative := x < 0

	if isNegative {
		x *= -1
	}

	chars := []byte(strconv.Itoa(x))

	left := 0
	right := len(chars) - 1

	for right > left {
		tmp := chars[right]
		chars[right] = chars[left]
		chars[left] = tmp

		left++
		right--
	}

	res, _ := strconv.Atoi(string(chars))
	if isNegative {
		res *= -1
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}*/
