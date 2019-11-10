package main

import (
	"fmt"
)

func main4() {
	nums := []int{1, 2, 3, 4, 5, 6}
	nums = rotate(nums, 2)
	fmt.Println(nums)
}

func rotate(nums []int, n int) []int {
	temp := make([]int, len(nums))
	for i := range nums {
		index := i + n
		if index >= len(nums) {
			index -= len(nums)
		}
		temp[i] = nums[index]
	}
	return temp
}
