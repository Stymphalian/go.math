package matrix

import (
	"math"
)

const (
	epsilon = 0.000000001
)

func closeEquals(a, b, eps float64) bool {
	if a > b {
		return ((a - b) < eps)
	} else {
		return ((b - a) < eps)
	}
}

func degToRad(a float64) float64 {
	return a * math.Pi / 180.0
}
func radToDeg(a float64) float64 {
	return a * 180.0 / math.Pi
}

func clamp(a, lower, upper float64) float64 {
	if a < lower {
		return lower
	} else if a > upper {
		return upper
	}
	return a
}
