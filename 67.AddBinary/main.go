package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}

	dif := len(a) - len(b)
	carry := "0"
	for i := 0; i < dif; i++ {
		b = carry + b
	}

	result := ""

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if carry == "1" {
				result = "1" + result
			} else {
				result = "0" + result
				carry = "1"
			}
		} else if a[i] == '0' && b[i] == '0' {
			if carry == "1" {
				result = "1" + result
				carry = "0"
			} else {
				result = "0" + result
			}
		} else if a[i] != b[i] {
			if carry == "1" {
				result = "0" + result
			} else {
				result = "1" + result
			}
		}
	}
	if carry == "1" {
		result = "1" + result
	}
	return result
}

func main() {
	a, b := "10110", "11"
	fmt.Println(a, b)
	fmt.Println(addBinary(a, b))
}
