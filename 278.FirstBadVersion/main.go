package main

import "fmt"

func isBadVersion(version int) bool {
	return version >= 4
}

func firstBadVersion(n int) int {
	pMin, pMax, res := 0, n, n/2
	for pMax > pMin {
		if isBadVersion(res) {
			pMax = res
			res = (pMax + pMin) / 2
		} else {
			pMin = res + 1
			res = (pMax + pMin) / 2
		}
	}

	return res
}

func main() {
	fmt.Println(firstBadVersion(10))
}
