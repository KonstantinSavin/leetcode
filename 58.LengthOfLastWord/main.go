package main

import "fmt"

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 && string(s[i]) == " " {
		i--
	}
	lastIndex := i
	for i >= 0 && string(s[i]) != " " {
		i--
	}
	return lastIndex - i
}

func main() {
	s := "Hello World"
	// fmt.Println(string(s[6]))
	fmt.Println(lengthOfLastWord(s))
}
