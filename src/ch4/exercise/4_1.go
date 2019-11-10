/*
4.1 编写一个函数，计算两个SHA256哈希码中不同bit的数目。(参考2.6.2节的PopCount 函数。)
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

func main1() {
	ch1 := sha256.Sum256([]byte("xc"))
	ch2 := sha256.Sum256([]byte("addd"))
	res := DiffCount(ch1, ch2)
	fmt.Printf("ch1 diff ch2 %d bit\n", res)
}

func DiffCount(ch1 [32]byte, ch2 [32]byte) int {
	res := 0
	for i := 0; i < 32; i++ {
		res += xor(ch1[i], ch2[i])
	}
	return res
}

func xor(n1 byte, n2 byte) int {
	temp := n1 ^ n2
	res := 0
	for i := 0; i < 8; i++ {
		if temp&(1<<i) != 0 {
			res += 1
		}
	}
	return res
}
