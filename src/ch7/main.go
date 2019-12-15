package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(io.Writer)

	if w, ok := w.(*bytes.Buffer); ok {
		fmt.Println(w)
	}
	f.Write([]byte("hello\n"))

}
