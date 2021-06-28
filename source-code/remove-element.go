package main

import (
	"fmt"
	"time"
)

func removeElement(nums []int, val int) int {
	var (
		i = 0
		j = 0
	)
	for i < len(nums) {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	if len(nums) > 0 {
		nums = nums[:j]
	}
	return len(nums)
}

func main() {
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeElementStart := time.Now()
	fmt.Println(removeElement(nums, 1))
	fmt.Println(time.Since(removeElementStart))
}
