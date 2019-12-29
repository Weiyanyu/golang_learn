package main

import (
	"fmt"
	"reflect"
)

type Moive struct {
	Title, SubTitle string
	Year            int
	Color           bool `json:"color"`
}

func main() {
	m := Moive{"aaa", "bbb", 2019, true}
	s := reflect.ValueOf(m).Type().Field(3).Tag.Get("json")
	fmt.Println(s)
}
