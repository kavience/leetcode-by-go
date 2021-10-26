package source_code

// 方法一
func TwoSum1(nums []int, target int) (result []int) {
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
func TwoSum2(nums []int, target int) []int {
	numMap := map[int]int{}
	for key, value := range nums {
		if k, ok := numMap[target-value]; ok {
			return []int{k, key}
		}
		numMap[value] = key
	}
	return nil
}
