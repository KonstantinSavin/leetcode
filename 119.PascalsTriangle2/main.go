package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	pascalsTriangle := [][]int{}
	numRows := rowIndex + 1
	pascalsTriangle = append(pascalsTriangle, []int{1})
	if numRows == 1 {
		return pascalsTriangle[0]
	}

	for i := 1; i < numRows; i++ {
		pascalsTriangle = append(pascalsTriangle, pascalTriangleMaker(pascalsTriangle[i-1]))
	}
	row := pascalsTriangle[rowIndex]

	return row
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
	fmt.Println(getRow(2))
}
