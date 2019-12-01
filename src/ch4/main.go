package main

import "fmt"

func main() {

	AddEdge("1", "3")
	AddEdge("2", "3")

	fmt.Printf("has 1 -> 3 ? %v\n", hasEdge("1", "3"))
	fmt.Printf("has 4 -> 3 ? %v\n", hasEdge("4", "3"))

}
