package main

import (
	"flag"
	"fmt"
	"go_learn/src/ch8/du"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var w sync.WaitGroup

	for _, root := range roots {
		w.Add(1)
		go du.Walk(root, &w, fileSizes)
	}

	go func() {
		w.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(100 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:

	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nbytes += size
			nfiles += 1
		case <-tick:
			fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
		}
	}

	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
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
