package main

import "fmt"

func main() {
	fmt.Println(addStrings("12", "123"))
}

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var result []byte

	for i >= 0 || j >= 0 || carry > 0 {
		var digit1, digit2 int

		if i >= 0 {
			digit1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			digit2 = int(num2[j] - '0')
			j--
		}

		sum := digit1 + digit2 + carry
		carry = sum / 10
		result = append(result, byte(sum%10)+'0')
	}

	for k, n := 0, len(result); k < n/2; k++ {
		result[k], result[n-1-k] = result[n-1-k], result[k]
	}

	return string(result)
}
