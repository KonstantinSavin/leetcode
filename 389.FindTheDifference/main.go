package main

import "fmt"

func findTheDifference(s string, t string) byte {
	h := [26]int{}

	for _, v := range t {
		h[v-'a']++
	}

	for _, v := range s {
		h[v-'a']--
	}

	for i, v := range h {
		if v > 0 {
			return byte(i + 'a')
		}
	}
	return 0
}

func main() {
	fmt.Println(findTheDifference("abcd", "abecd"))

}
