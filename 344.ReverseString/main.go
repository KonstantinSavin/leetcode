package main

import "fmt"

func reverseString(s []byte) []byte {
	for l, r := 0, len(s)-1; l < r; {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}

func main() {
	fmt.Println(reverseString([]byte{'a', 'b', 'e', 'i', 'o', 'u', 'n'}))
}
