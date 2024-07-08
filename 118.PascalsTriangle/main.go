package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	result := [][]int{}
	if numRows == 0 {
		return result
	}

	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}

	for i := 1; i < numRows; i++ {
		result = append(result, pascalTriangleMaker(result[i-1]))
	}

	return result
}

func pascalTriangleMaker(p []int) []int {
	res := make([]int, 1, len(p)+1)
	res = append(res, p...)

	for i := 0; i < len(res)-1; i++ {
		res[i] += res[i+1]
	}

	return res
}

func main() {
	fmt.Println(generate(3))
}
