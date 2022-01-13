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

func SqrtN(m float32, n float32) float32 {
	m64 := float64(m)
	n64 := float64(n)
	return float32(SqrtN64(m64, n64))
}

// 开N次方
func SqrtN64(m float64, n float64) float64 {
	z := float64(1)
	tmp := float64(0)
	for math.Abs(tmp-z) > 0.0000000001 {
		tmp = z
		z = z - (math.Pow(z, n)-m)/(n*math.Pow(z, n-1))
	}
	return z
}
