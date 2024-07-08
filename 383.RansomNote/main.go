package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	ht := make(map[rune]int)
	for _, k := range magazine {
		ht[k-'a']++
	}
	for _, k := range ransomNote {
		ht[k-'a']--
		if ht[k-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("abd", "aabccc"))
}
