package unit

import (
	"testing"
)

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
