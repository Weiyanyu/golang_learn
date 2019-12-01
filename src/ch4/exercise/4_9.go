package main

import "bufio"

import "os"

import "fmt"

func main() {
	counts := make(map[string]int)
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Split(bufio.ScanWords)

	for scaner.Scan() {
		counts[scaner.Text()]++
	}

	fmt.Printf("word\tsize\n")
	for word, size := range counts {
		fmt.Printf("%s\t%d\n", word, size)
	}

}
