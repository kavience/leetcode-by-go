package source_code

func RemoveDuplicates(nums []int) int {
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
