package main

import "fmt"

func isZero(d [26]int) bool {
	for i := 0; i < 26; i++ {
		if d[i] != 0 {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	d := [26]int{}
	for i := 0; i < len(s1); i++ {
		d[s1[i]-'a']++
		d[s2[i]-'a']--
	}
	if isZero(d) {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		d[s2[i-len(s1)]-'a']++
		d[s2[i]-'a']--
		if isZero(d) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("adc", "dcda"))
}
