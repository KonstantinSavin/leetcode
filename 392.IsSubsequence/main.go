package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	var p int
	for _, v := range t {
		if p == len(s) {
			return true
		} else if s[p] == byte(v) {
			p++
		}
	}
	return p == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "abedce"))

}
