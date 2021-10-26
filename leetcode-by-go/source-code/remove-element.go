package source_code

func RemoveElement(nums []int, val int) int {
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
