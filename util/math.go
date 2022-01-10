package util

import "math"

// 绝对值
func Abs(i float32) float32 {
	if i < 0 {
		i = 0 - i
	}
	return i
}

// 最大值
func MaxInt(values ...int) int {
	var max int
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

// 最大值
func Max(values ...float32) float32 {
	var max float32
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

// 最小值
func Min(values ...float32) float32 {
	var min float32
	for i, value := range values {
		if i == 0 {
			min = value
		} else if value < min {
			min = value
		}
	}
	return min
}

// 向下取整
func Floor(value float32) float32 {
	return float32(math.Floor(float64(value)))
}
