package main

import "fmt"

func addDigits(num int) int {
	for num > 9 {
		res := 0
		for num > 0 {
			res += num % 10
			num /= 10
		}
		num = res
	}

	return num
}

func addDigits2(num int) int {
	return num - 9*((num-1)/9)
}

func main() {
	fmt.Println(addDigits(5), addDigits(38))
	fmt.Println(addDigits2(5), addDigits2(38))
}
