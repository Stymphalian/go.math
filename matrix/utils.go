package matrix

import (
	"math"
)

const (
	epsilon = 0.000000001
)

// Checks if two floats are equal. Doing a comparision
// using a small epsilon value
func closeEquals(a, b, eps float64) bool {
	if a > b {
		return ((a - b) < eps)
	} else {
		return ((b - a) < eps)
	}
}


// Convert from a degree to a radian
func degToRad(a float64) float64 {
	return a * math.Pi / 180.0
}

// Convert radian to degree
func radToDeg(a float64) float64 {
	return a * 180.0 / math.Pi
}

// Clamp the a value between the lower and upper
// Therefore the value returned is between the range [lower,upper] (inclusive)
func clamp(a, lower, upper float64) float64 {
	if a < lower {
		return lower
	} else if a > upper {
		return upper
	}
	return a
}
