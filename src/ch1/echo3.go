package main

import (
	"fmt"
	"os"
	"strings"
)

func main7() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}