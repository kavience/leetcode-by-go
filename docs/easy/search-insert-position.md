# 搜索插入位置 

难度：1/10

类型：数组

## 题目
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

> 输入: [1,3,5,6], 5
> 
> 输出: 2

示例 2:

> 输入: [1,3,5,6], 2
> 
> 输出: 1

示例 3:

> 输入: [1,3,5,6], 7
> 
> 输出: 4

## 解法

```go
func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		} else {
			if target > nums[i] && target < nums[i+1] {
				return i + 1
			}
		}
	}
	return 0
}
```


> 来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-insert-position
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。