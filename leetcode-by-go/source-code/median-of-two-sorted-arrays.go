package source_code

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}
	left, right, m, n, median1, median2 := 0, len(nums1), len(nums1), len(nums2), 0, 0
	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i
		nums_im1 := math.MinInt32
		nums_i := math.MaxInt32
		if i != 0 {
			nums_im1 = nums1[i-1]
		}
		if i != m {
			nums_i = nums1[i]
		}
		nums_jm1 := math.MinInt32
		nums_j := math.MaxInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		if j != n {
			nums_j = nums2[j]
		}
		if nums_im1 <= nums_j {
			median1 = int(math.Max(float64(nums_im1), float64(nums_jm1)))
			median2 = int(math.Min(float64(nums_i), float64(nums_j)))
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64((median1 + median2)) / 2.0
	}
	return float64(median1)
}
