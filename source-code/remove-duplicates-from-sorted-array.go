package main

import (
	"fmt"
	"time"
)

func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[j] != nums[i+1] {
			nums[j+1] = nums[i+1]
			j++
		}
	}
	if len(nums) > 0 {
		nums = nums[:j+1]
	}
	return len(nums)
}

func main() {
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicatesStart := time.Now()
	fmt.Println(removeDuplicates(nums))
	fmt.Println(time.Since(removeDuplicatesStart))
}
