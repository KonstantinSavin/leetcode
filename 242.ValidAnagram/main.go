package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var hashArr [26]int
	for i := 0; i < len(s); i++ {
		hashArr[s[i]-'a']++
		hashArr[t[i]-'a']--
	}

	for i := range hashArr {
		if hashArr[i] != 0 {
			return false
		}
	}
	fmt.Println(hashArr)
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("anagrama", "nagaram"))
}
