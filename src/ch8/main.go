package main

import (
	"flag"
	"go_learn/src/ch8/chat"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	chat.TestChat()
}

var tokens = make(chan struct{}, 20)

func testDemo(str string) []string {
	tokens <- struct{}{}
	var res []string
	switch str {
	case "1":
		res = []string{"2", "3"}
	case "2":
		res = []string{"1", "3"}
	case "3":
		res = []string{"1", "2"}
	default:
		res = []string{}
	}
	<-tokens
	return res
}
