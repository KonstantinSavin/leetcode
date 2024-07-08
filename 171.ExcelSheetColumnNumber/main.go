package main

import "fmt"

func titleToNumber(columnTitle string) int {
	res, mul := 0, 1
	len := len(columnTitle)
	if len > 0 {
		for i := len - 1; i >= 0; i-- {
			res += int(columnTitle[i]-'A'+1) * mul
			mul *= 26
		}
	}

	return res
}
func main() {
	fmt.Println(titleToNumber("A"), titleToNumber("C"), titleToNumber("AA"))
	fmt.Println(int('A'))
}
