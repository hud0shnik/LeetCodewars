package main

import "fmt"

func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := "" // int: 1994 - 1000 -> 994 - 900 -> 94 -90 -> 4 - 4
	i := 0       // result: "" -> ""+"M" -> "M"+"CM" -> "MCM"+"XC" -> "MCMXC"+ "IV"
	for num > 0 {
		for romanValues[i] > num {
			i++
		}
		result += roman[i]
		num -= romanValues[i]
	}
	return result
}

func main() {
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1))
}
