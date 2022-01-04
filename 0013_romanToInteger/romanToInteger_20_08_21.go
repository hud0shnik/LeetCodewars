package main

import (
	"fmt"
	"time"
)

func romanToInt(s string) int {
	result := 0
	numMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	i := 0

	for ; i < len(s); i++ {
		char := string(s[i])
		if char == "I" {
			if i < len(s)-1 {
				next := string(s[i+1])
				if next == "V" || next == "X" {
					result += numMap[next] - numMap[char]
					i++
					continue
				}
			}
		}
		if char == "X" {
			if i < len(s)-1 {
				next := string(s[i+1])
				if next == "L" || next == "C" {
					result += numMap[next] - numMap[char]
					i++
					continue
				}
			}
		}
		if char == "C" {
			if i < len(s)-1 {
				next := string(s[i+1])
				if next == "D" || next == "M" {
					result += numMap[next] - numMap[char]
					i++
					continue
				}
			}
		}
		result += numMap[char]
	}
	return result
}

func main() {
	start := time.Now()
	//	fmt.Println(romanToInt("LVIII"))
	for k := 0; k < 10000; k++ {
		fmt.Println(romanToInt("MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMDCXLVII") + k)
	}
	for k := 0; k < 10000; k++ {
		fmt.Println(romanToInt("MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMDCXLVII") * k)
	}
	for k := 0; k < 10000; k++ {
		fmt.Println(romanToInt("MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMDCXLVII") ^ k)
	}
	//fmt.Println(romanToInt("MCMXCIV"))

	duration := time.Since(start)
	fmt.Println("\ttime:", duration /*.Microseconds()*/)
}
