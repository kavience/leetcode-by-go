# 盛最多水的容器

难度：4/10

类型：数组

## 题目
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

`说明`：你不能倾斜容器。

示例 1：

<img :src="$withBase('/container-with-most-water.jpeg')" alt="container-with-most-water">

> 输入：[1,8,6,2,5,4,8,3,7]
> 
> 输出：49 
> 
> 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。


## 解法一
暴力循环，两层 for 循环

时间复杂度：O(n^2)

```go
// 方法一
func maxArea1(height []int) int {
	var result = 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			x := j - i
			y := math.Min(float64(height[i]), float64(height[j]))
			area := x * int(y)
			if result < area {
				result = area
			}
		}
	}
	return result
}
```

## 解法二

动态规划，左右夹逼，只循环一次

时间复杂度：O(n)

```go
// 方法二
func maxArea2(height []int) int {
	var (
		result = 0
		i      = 0
		j      = len(height) - 1
	)
	for i < j {
		var y = math.Min(float64(height[i]), float64(height[j]))
		result = int(math.Max(float64(result), float64(int(y)*(j-i))))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}
```

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/container-with-most-water

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。