/*
重写reverse函数，使用数组指针代替slice。
*/
package main

import (
	"fmt"
)

func main2() {
	nums := [6]int{1, 2, 3, 4, 5, 6}
	reverse(&nums)
	fmt.Println(nums)
}

func reverse(nums *[6]int) {
	for i, j := 0, len(*nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
