package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
}

func repeatedSubstringPattern(s string) bool {
	if s == "" {
		return false
	}

	n := len(s)

	for l := 1; l <= n/2; l++ {
		if n%l == 0 {
			ss := s[:l]

			if strings.Repeat(ss, n/l) == s {
				return true
			}
		}
	}
	return false
}
