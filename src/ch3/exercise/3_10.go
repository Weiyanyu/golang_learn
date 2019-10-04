/*
练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
 */
package main

import (
	"bytes"
	"fmt"
)

func main() {

	fmt.Println(comma1("1231231231233"))
}

func comma1(s string) string {
	var buf bytes.Buffer
	n := len(s)
	count := n % 3
	if count == 0 {
		count = 3
	}
	for i := 0; i < n; {
		if count == 0 {
			buf.WriteByte(',')
			count = 3
		} else {
			buf.WriteByte(s[i])
			i++
			count--
		}
	}
	return buf.String()
}


