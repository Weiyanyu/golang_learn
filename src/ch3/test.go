package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = 1 << (10 * iota)
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Tuesday)
}
