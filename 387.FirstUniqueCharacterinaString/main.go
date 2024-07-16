package main

import "fmt"

func firstUniqChar(s string) int {
	var hash [26]int
	for _, i := range s {
		hash[i-'a']++
	}

	for i := 0; i < len(s); i++ {
		if hash[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {

	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
