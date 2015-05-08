package lmath

import (
	"math"
)

const (
	epsilon = 0.000000001
)

// Checks if two floats are equal. Doing a comparision
// using a small epsilon value
func closeEq(a, b, eps float64) bool {
	if a > b {
		return ((a - b) < eps)
	} else {
		return ((b - a) < eps)
	}
}

// Convert from a degree to a radian
func Radians(a float64) float64 {
	return a * math.Pi / 180.0
}

// Convert radian to degree
func Degrees(a float64) float64 {
	return a * 180.0 / math.Pi
}

// Clamp the a value between the lower and upper
// Therefore the value returned is between the range [lower,upper] (inclusive)
func Clamp(a, lower, upper float64) float64 {
	if a < lower {
		return lower
	} else if a > upper {
		return upper
	}
	return a
}

// Lineat interpolates between the start and end values
// inc is specified between the range 0 -1
// 0  --> start
// x -->  somewhere between start and end
// 1 --> end
func Lerp(start, end, inc float64) float64 {
	return (1-inc)*start + inc*end
}
