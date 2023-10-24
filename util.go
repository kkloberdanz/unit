package unit

import (
	"math"
)

const (
	float64EqualityThreshold = 1e-16
)

func almostEqual(a, b float64) bool {
	return math.Abs(float64(a-b)) <= float64EqualityThreshold
}
