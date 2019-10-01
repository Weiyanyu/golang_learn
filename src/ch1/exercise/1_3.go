/*
练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
 */

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main112() {
	start := time.Now().Nanosecond()
	lowEcho()
	end := time.Now().Nanosecond()
	fmt.Printf("low Echo version need %d ns\n", end - start)

	start = time.Now().Nanosecond()
	coolEcho()
	end = time.Now().Nanosecond()
	fmt.Printf("cool Echo version need %d ns\n", end - start)
}

func lowEcho() {
	s, sep := "", ""

	for _, arg := range os.Args[1: ] {
		s += sep + arg
	}
}

func coolEcho() {
	strings.Join(os.Args[1:], " ")
}
