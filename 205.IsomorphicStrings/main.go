package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sliceS := make([]int, 256)
	sliceT := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if sliceS[int(s[i])] != sliceT[int(t[i])] {
			return false
		}

		i2 := i + 1
		sliceS[int(s[i])] = i2
		sliceT[int(t[i])] = i2
	}

	return true
}

func main() {
	s := "adde"
	t := "egge"
	fmt.Println(isIsomorphic(s, t))
}
