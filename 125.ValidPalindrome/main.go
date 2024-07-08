package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	clean_s := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
			clean_s = append(clean_s, s[i])
		}
	}

	fmt.Println(clean_s)
	for i := 0; i < len(clean_s); i++ {
		if clean_s[i] != clean_s[len(clean_s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))

}
