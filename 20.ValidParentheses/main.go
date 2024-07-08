package main

import "fmt"

func isValid(s string) bool {
	leftSymbol := make([]string, 0)
	for i := range s {
		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			leftSymbol = append(leftSymbol, string(s[i]))
		} else if string(s[i]) == ")" && len(leftSymbol) != 0 && leftSymbol[len(leftSymbol)-1] == "(" {
			leftSymbol = leftSymbol[:len(leftSymbol)-1]
		} else if string(s[i]) == "]" && len(leftSymbol) != 0 && leftSymbol[len(leftSymbol)-1] == "[" {
			leftSymbol = leftSymbol[:len(leftSymbol)-1]
		} else if string(s[i]) == "}" && len(leftSymbol) != 0 && leftSymbol[len(leftSymbol)-1] == "{" {
			leftSymbol = leftSymbol[:len(leftSymbol)-1]
		} else {
			return false
		}
		fmt.Println(string(s[i]))
	}
	fmt.Println(leftSymbol, len(leftSymbol))
	return len(leftSymbol) == 0
}

func main() {

	fmt.Println(isValid("()[]{}"))
}
