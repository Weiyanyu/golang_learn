package main

import (
	"fmt"
	"go_learn/src/ch9/memo"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var bookPrice map[string]float64
var loadOnce sync.Once

func main() {
	urls := []string{
		"https://www.baidu.com/",
		"https://www.iqiyi.com/",
		"https://studygolang.com/",
		"https://www.iqiyi.com/",
		"https://tech.meituan.com/",
		"https://www.baidu.com/",
		"https://tech.meituan.com/",
		"https://www.baidu.com/",
	}

	m := memo.New(httpGetBody)
	var w sync.WaitGroup
	for _, url := range urls {
		w.Add(1)
		go func(url string) {
			// start := time.Now()
			_, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			// fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			w.Done()
		}(url)
	}

	w.Wait()
}

func BookPrice(bookName string, ch chan<- float64, w *sync.WaitGroup) {
	defer w.Done()
	loadOnce.Do(loadBookPrice)
	ch <- bookPrice[bookName]
}

func loadBookPrice() {
	bookPrice = map[string]float64{
		"C++":    1.99,
		"Java":   2.99,
		"golang": 0.99,
	}
}

func httpGetBody(url string) (interface{}, error) {
	fmt.Println("httpGetBody被调用")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
