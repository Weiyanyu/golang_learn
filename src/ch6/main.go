package main

import "fmt"

func main() {
	path := Path{
		{1, 1},
		{4, 6},
		{8, 4},
	}
	offset := Point{2, 2}

	fmt.Println(path.TranslateBy(offset, Point.Add))
}
