# 三数之和 

难度：3/10

类型：数组

## 题目
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

> 输入：nums = [-1,0,1,2,-1,-4]
> 
> 输出：[[-1,-1,2],[-1,0,1]]

示例 2：

> 输入：nums = []
> 
> 输出：[]

## 解法

```go
func threeSum(nums []int) [][]int {
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
```

> 来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。