package main

import "fmt"

func main() {
	// double(4)
	err := Recover()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
