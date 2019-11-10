package main

import (
	"fmt"
)

func main() {
	a := []int{}

	a = AppendInt(a, 1)
	a = AppendInt(a, 2)
	fmt.Println(a)

}
