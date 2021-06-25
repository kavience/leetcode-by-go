package main

import (
	"fmt"
	"time"
)

// 方法一
func twoSum1(nums []int, target int) (result []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				return
			}
		}
	}
	return
}

// 方法二
func twoSum2(nums []int, target int) []int {
	numMap := map[int]int{}
	for key, value := range nums {
		if k, ok := numMap[target-value]; ok {
			return []int{k, key}
		}
		numMap[value] = key
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4, 6, 8, 10}
	target := 12
	twoSum1Start := time.Now()
	fmt.Println(twoSum1(nums, target))
	fmt.Println(time.Since(twoSum1Start))

	twoSum2Start := time.Now()
	fmt.Println(twoSum2(nums, target))
	fmt.Println(time.Since(twoSum2Start))
}
