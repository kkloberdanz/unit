package temperature

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-16

func almostEqualC(a, b Celcius) bool {
	return math.Abs(float64(a-b)) <= float64EqualityThreshold
}

func almostEqualF(a, b Farenheit) bool {
	return math.Abs(float64(a-b)) <= float64EqualityThreshold
}

func TestCelciusToFarenheit(t *testing.T) {
	tests := []struct {
		input    Celcius
		expected Farenheit
	}{
		{0.0, 32.0},
		{100.0, 212.0},
	}

	for _, test := range tests {
		result := CelciusToFarenheit(test.input)
		if !almostEqualF(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}

func TestFarenheitToCelcius(t *testing.T) {
	tests := []struct {
		input    Farenheit
		expected Celcius
	}{
		{0.0, -17.77777777777777856727},
		{100.0, 37.77777777777777856727},
	}

	for _, test := range tests {
		result := FarenheitToCelcius(test.input)
		if !almostEqualC(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}
