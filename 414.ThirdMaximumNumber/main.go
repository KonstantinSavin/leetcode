package main

import "fmt"

func main() {
	numbers := []int{3, 5, 7, 2, 8, -1, 4, 10, 2}
	fmt.Println(thirdMax(numbers))
	fmt.Println(thirdMax([]int{2}))
	fmt.Println(thirdMax([]int{3, 1}))
	fmt.Println(thirdMax([]int{3, 4}))
}

func thirdMax(nums []int) int {
	uniqueValues := make(map[int]struct{})
	for _, value := range nums {
		uniqueValues[value] = struct{}{}
	}

	uniqueSlice := make([]int, 0, len(uniqueValues))
	for key := range uniqueValues {
		uniqueSlice = append(uniqueSlice, key)
	}

	var res int

	if len(uniqueSlice) == 1 {
		res = uniqueSlice[0]
	} else if len(uniqueSlice) == 2 {
		res = uniqueSlice[0]
		if res < uniqueSlice[1] {
			res = uniqueSlice[1]
		}
	} else {
		const minInt = -1 << 31

		firstMax, secondMax, thirdMax := minInt, minInt, minInt

		for _, num := range uniqueSlice {
			if num == firstMax || num == secondMax || num == thirdMax {
				continue
			}
			if num > firstMax {
				thirdMax = secondMax
				secondMax = firstMax
				firstMax = num
			} else if num > secondMax {
				thirdMax = secondMax
				secondMax = num
			} else if num > thirdMax {
				thirdMax = num
			}
		}
		res = thirdMax
	}

	return res
}
