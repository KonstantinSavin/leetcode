package main

import "fmt"

func romanToInt(s string) int {
	romanToIntDic := make(map[string]int)
	romanToIntDic["I"] = 1
	romanToIntDic["V"] = 5
	romanToIntDic["X"] = 10
	romanToIntDic["L"] = 50
	romanToIntDic["C"] = 100
	romanToIntDic["D"] = 500
	romanToIntDic["M"] = 1000
	fmt.Println(romanToIntDic)

	var result = 0
	for i := len(s) - 1; i >= 0; i-- {
		num := romanToIntDic[string(s[i])]
		if 4*num < result {
			result -= num
		} else {
			result += num
		}
		fmt.Println(result, num)
	}
	return result
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
