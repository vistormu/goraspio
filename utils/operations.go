package utils

import (
    "math"
)


type Number interface {
    ~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Clip[T Number](value, min, max T) T {
    if value < min {
        return min
    }
    if value > max {
        return max
    }
    return value
}

func Mean[T Number](values []T) T {
    var sum T
    length := len(values)

    for _, v := range values {
        sum += v
    }

    return sum / T(length)
}

func StdDev[T Number](values []T) T {
	if len(values) == 0 {
		return 0
	}

	var sum T
	for _, v := range values {
		sum += v
	}

	mean := sum / T(len(values))

	var varianceSum T
	for _, v := range values {
		diff := v - mean
		varianceSum += diff * diff
	}

	variance := varianceSum / T(len(values))
	return T(math.Sqrt(float64(variance)))
}