package main

import "log"

func square(n int) int {
	return n * n
}

func cube(s func(int) int, n int) int {
	return s(n) * n
}

func testCallback(str string) int {
	if str == "" {
		return 0
	}
	return 1
}

func bf(str string, callback func(string) int) {
	n := callback(str)
	if n == 0 {
		log.Println("has some error")
	} else {
		log.Println("good")
	}
}

func main() {

	bf("", testCallback)
	bf("hello, world", testCallback)
}
