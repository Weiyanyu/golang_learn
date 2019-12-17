package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func Walk(dir string, w *sync.WaitGroup, fileSizes chan<- int64) {
	defer w.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			w.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go Walk(subDir, w, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

var tokens = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	tokens <- struct{}{}
	defer func() { <-tokens }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}
