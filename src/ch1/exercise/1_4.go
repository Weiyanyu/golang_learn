/*
练习 1.4： 修改dup2 ，出现重复的行时打印文件名称。
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main33() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLine(os.Stdin, counts)
	} else {
		for _, filename := range filenames {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLine(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] > 0 {
			fmt.Printf("repeat filename : %s\n", f.Name())
		}
		counts[line]++
	}
}
