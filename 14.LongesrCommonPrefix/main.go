package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	lenStrs := len(strs)

	if lenStrs == 0 {
		return ""
	}

	var prefix = ""

	fmt.Println(strs[0])
	for i := 0; i < len(strs[0]); i++ {
		letter := string(strs[0][i])
		fmt.Println(letter)
		match := true
		for j := 1; j < len(strs); j++ {
			if (len(strs[j]) - 1) < i {
				match = false
				break
			}

			if string(strs[j][i]) != letter {
				match = false
				break
			}

		}
		if match {
			prefix += letter
		} else {
			break
		}
	}

	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
