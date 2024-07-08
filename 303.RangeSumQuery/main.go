package main

import "fmt"

type NumArray struct {
	sumArr []int
}

func Constructor(nums []int) NumArray {
	sumArr := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sumArr[i] = sumArr[i-1] + nums[i-1]
	}
	return NumArray{sumArr: sumArr}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sumArr[right+1] - this.sumArr[left]
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(Constructor(nums))
}
