package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			dif := prices[i] - min
			if dif > res {
				res = dif
			}
		}
	}

	return res
}

func main() {
	prices := []int{7, 2, 5}
	fmt.Println(maxProfit(prices))
}
