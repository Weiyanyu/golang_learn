/*
练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("response status code is : %s", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s : %v\n", url ,err)
			os.Exit(1)
		}
	}

}
