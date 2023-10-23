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
		{12312.0, 22193.60000000000218278728},
		{0.0000043, 32.00000774000000092201},
		{1234.0, 2253.20000000000027284841},
		{659846935.0, 1187724515.00000000000000000000},
		{1054004.0, 1897239.19999999995343387127},
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
		{12312.0, 6822.22222222222262644209},
		{0.0000043, -17.77777538888889097279},
		{1234.0, 667.77777777777782830526},
		{659846935.0, 366581612.77777779102325439453},
		{1054004.0, 585540.00000000000000000000},
	}

	for _, test := range tests {
		result := FarenheitToCelcius(test.input)
		if !almostEqualC(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}
