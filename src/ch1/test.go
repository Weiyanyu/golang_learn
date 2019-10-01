package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("0.1")
	if err != nil {
		n = 5
	}
	fmt.Println(n)
}
