package main

import (
	"fmt"
)

func main() {
	str := []string{"a", "a", "b", "b", "d", "b", "c", "c"}
	fmt.Println(RemoveDuplicates(str))
}

func RemoveDuplicates(strs []string) []string {
	for i := 0; i < len(strs)-1; {
		if strs[i] == strs[i+1] {
			copy(strs[i:], strs[i+1:])
			strs = strs[:len(strs)-1]
		} else {
			i++
		}
	}
	return strs
}
