# 两数之和

难度：1/10

类型：数组

## 题目
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

示例 1：

> 输入：nums = [2,7,11,15], target = 9
> 
> 输出：[0,1]
> 
> 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：

> 输入：nums = [3,2,4], target = 6
> 
> 输出：[1,2] 

示例 3：

> 输入：nums = [3,3], target = 6
> 
> 输出：[0,1]


## 解法一
暴力循环，两层 for 循环

时间复杂度：O(n^2)

```go
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
```

## 解法二

利用 map 的特性

时间复杂度：O(n)

```go
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
```

> 来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
