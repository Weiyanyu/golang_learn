package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.CtoK(tempconv.FreezingC))
}