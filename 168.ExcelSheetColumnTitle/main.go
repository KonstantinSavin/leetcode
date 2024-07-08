package main

import "fmt"

func convertToTitle(columnNumber int) string {
	res := ""

	for columnNumber > 0 {
		columnNumber--
		res = string(byte(columnNumber%26)+'A') + res
		columnNumber /= 26
	}

	return res
}

func main() {
	fmt.Println(convertToTitle(58))
}
