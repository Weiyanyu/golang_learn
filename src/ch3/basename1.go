package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("abc/c.go"))
	fmt.Println(basename2("abc/c"))

}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[0:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	index := strings.LastIndex(s, "/")
	s = s[index+1:]
	if index = strings.LastIndex(s, ","); index >= 0 {
		s = s[0: index]
	}
	return s
}
