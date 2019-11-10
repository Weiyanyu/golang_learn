package main

import (
	"crypto/sha256"
	"fmt"
)

func main1() {
	ch1 := sha256.Sum256([]byte("x112"))
	ch2 := sha256.Sum256([]byte("X"))

	fmt.Println(ch1 == ch2)
}
