package source_code

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	// 先排序
	sort.Ints(nums)
	a, b, c := 0, 0, 0

	for ; a < n; a++ {
		// 跳过重复的第一个值
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		// 每次循环把第三个值重置为最后一个
		c = n - 1
		// 循环第二个值
		for b = a + 1; b < n; b++ {
			// 跳过重复的第二个值
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			// 递减第三个值，直到第二个值大于或等于第三个值或者三者之和大于0
			for b < c && nums[b]+nums[c]+nums[a] > 0 {
				c--
			}
			// 如果第二个值与第三个相等，则没有找到
			if b == c {
				break
			}
			// 如果三者之和等于0，说明找到了
			if nums[b]+nums[c]+nums[a] == 0 {
				result = append(result, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return result
}
