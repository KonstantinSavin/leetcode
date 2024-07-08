package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	sl := strings.Split(s, "")
	fmt.Println(sl)
	for l, r := 0, len(s)-1; l < r; {
		for l < r && !strings.ContainsAny(sl[l], vowels) {
			l++
		}
		for l < r && !strings.ContainsAny(sl[r], vowels) {
			r--
		}
		sl[l], sl[r] = sl[r], sl[l]
		l++
		r--
	}
	s = strings.Join(sl, "")
	return s
}

func main() {
	fmt.Println(reverseVowels("abeioun"))
}
