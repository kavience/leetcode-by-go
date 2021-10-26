package source_code

import (
	"math"
)

// 方法一
func MaxArea1(height []int) int {
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

// 方法二
func MaxArea2(height []int) int {
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
