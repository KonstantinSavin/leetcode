package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
}

func countSegments(s string) int {
	return len(strings.Fields(s))
}
