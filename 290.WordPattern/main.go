package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	var hashArr [26]string
	sArr := strings.Fields(s)

	if len(sArr) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if hashArr[pattern[i]-'a'] == "" {
			hashArr[pattern[i]-'a'] = sArr[i]
		} else if hashArr[pattern[i]-'a'] != sArr[i] {
			return false
		}
	}
	fmt.Println(hashArr)

	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			if hashArr[i] == hashArr[j] && hashArr[i] != "" {
				return false
			}
		}

	}

	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
	s = "dog cat cat fish"
	fmt.Println(wordPattern(pattern, s))
}
