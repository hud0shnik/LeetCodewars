package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	if len(s) < 2 {
		return s
	}
	dict := [256]int{}
	runeS := []rune(s)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	sort.Slice(runeS, func(i, j int) bool {
		if dict[runeS[i]] == dict[runeS[j]] {
			return runeS[i] > runeS[j]
		}
		return dict[runeS[i]] > dict[runeS[j]]
	})

	return string(runeS)
}

func main() {
	fmt.Println(frequencySort("tree"))
}
