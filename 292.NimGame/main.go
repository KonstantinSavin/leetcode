package main

import "fmt"

func canWinNim(n int) bool {
	if n%4 != 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(canWinNim(4))
	fmt.Println(canWinNim(1))
	fmt.Println(canWinNim(5))
}
