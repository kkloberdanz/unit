package unit

import (
	"testing"
)

func TestMeterToMile(t *testing.T) {
	tests := []struct {
		input    Meter
		expected Mile
	}{
		{1000, 0.6213712},
		{Meter(Kilo), 0.6213712},
		{Meter(Mega), 621.37119999999993069650},
	}

	for _, test := range tests {
		result := MeterToMile(test.input)
		if !almostEqual(float64(result), float64(test.expected)) {
			t.Errorf(
				"expected %.20f, but got %.20f",
				test.expected,
				result,
			)
		}
	}
}
