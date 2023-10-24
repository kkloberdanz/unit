package prefix

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-16

func almostEqual(a, b float64) bool {
	return math.Abs(float64(a-b)) <= float64EqualityThreshold
}

func TestScale(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{Yobibyte, 1208925819614629174706176},
	}

	for _, test := range tests {
		if !almostEqual(test.input, test.expected) {
			t.Errorf(
				"expected %.20f, but got %.20f",
				test.expected,
				test.input,
			)
		}
	}
}
