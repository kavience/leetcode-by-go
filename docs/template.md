# 标题 

难度：1/10

类型：数组

## 题目

## 解法

```go
func main() {
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicatesStart := time.Now()
	fmt.Println(removeDuplicates(nums))
	fmt.Println(time.Since(removeDuplicatesStart))
}
```