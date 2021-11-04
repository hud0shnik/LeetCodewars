package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	runeS, runeT := []rune(s), []rune(t)

	sort.Slice(runeS, func(i, j int) bool {
		return runeS[i] < runeS[j]
	})
	sort.Slice(runeT, func(i, j int) bool {
		return runeT[i] < runeT[j]
	})

	return string(runeS) == string(runeT)
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("car", "rat"))
	fmt.Println(isAnagram("uerirbuereghjcvbueregfreghjffrerijcvbhjcviyuguerwyuwegfegfreghyucvbueriyuwegweghjcvbriyuweiyuwegfreghjcvbueg", "reghjcvbueriyugewfreghjcvbueriyuwegfrjeghjcvbueriyuwegfreghjcvbueriyuwegfreghcvbueriyuwegfreghjcvbueriyuwegf"))
}
