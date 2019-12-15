package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan string)
	go func() {
		for _, filename := range os.Args[1:] {
			ch <- filename
		}
		close(ch)
	}()

	totalSize := MakeThnmbnails6(ch)
	fmt.Println(totalSize)
}
