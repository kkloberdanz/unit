package temperature

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-16

func almostEqualC(a, b Celcius) bool {
	return math.Abs(float64(a-b)) <= float64EqualityThreshold
}

func almostEqualF(a, b Fahrenheit) bool {
	return math.Abs(float64(a-b)) <= float64EqualityThreshold
}

func almostEqualK(a, b Kelvin) bool {
	return math.Abs(float64(a-b)) <= float64EqualityThreshold
}

func TestCelciusToFahrenheit(t *testing.T) {
	tests := []struct {
		input    Celcius
		expected Fahrenheit
	}{
		{0.0, 32.0},
		{100.0, 212.0},
		{273.15, 523.66999999999995907274},
		{12312.0, 22193.60000000000218278728},
		{0.0000043, 32.00000774000000092201},
		{1234.0, 2253.20000000000027284841},
		{659846935.0, 1187724515.00000000000000000000},
		{1054004.0, 1897239.19999999995343387127},
		{-100.432423, -148.77836139999999431893},
	}

	for _, test := range tests {
		result := CelciusToFahrenheit(test.input)
		if !almostEqualF(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}

func TestFahrenheitToCelcius(t *testing.T) {
	tests := []struct {
		input    Fahrenheit
		expected Celcius
	}{
		{0.0, -17.77777777777777856727},
		{100.0, 37.77777777777777856727},
		{273.15, 133.97222222222222853816},
		{12312.0, 6822.22222222222262644209},
		{0.0000043, -17.77777538888889097279},
		{1234.0, 667.77777777777782830526},
		{659846935.0, 366581612.77777779102325439453},
		{1054004.0, 585540.00000000000000000000},
		{-100.432423, -73.57356833333334122926},
	}

	for _, test := range tests {
		result := FahrenheitToCelcius(test.input)
		if !almostEqualC(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}

func TestKelvinToCelcius(t *testing.T) {
	tests := []struct {
		input    Kelvin
		expected Celcius
	}{
		{0.0, -273.15},
		{100.0, -173.14999999999997726263},
		{273.15, 0.0},
		{12312.0, 12038.85000000000036379788},
		{0.0000043, -273.14999569999997675041},
		{1234.0, 960.85000000000002273737},
		{659846935.0, 659846661.85000002384185791016},
		{1054004.0, 1053730.85000000009313225746},
	}

	for _, test := range tests {
		result := KelvinToCelcius(test.input)
		if !almostEqualC(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}

func TestCelciusToKelvin(t *testing.T) {
	tests := []struct {
		input    Celcius
		expected Kelvin
	}{
		{0.0, 273.15},
		{100.0, 373.14999999999997726263},
		{273.15, 546.29999999999995452526},
		{12312.0, 12585.14999999999963620212},
		{0.0000043, 273.15000429999997777486},
		{1234.0, 1507.15000000000009094947},
		{659846935.0, 659847208.14999997615814208984},
		{1054004.0, 1054277.14999999990686774254},
	}

	for _, test := range tests {
		result := CelciusToKelvin(test.input)
		if !almostEqualK(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	tests := []struct {
		input    Fahrenheit
		expected Kelvin
	}{
		{0.0, 255.37222222222220580079},
		{100.0, 310.92777777777774872447},
		{273.15, 407.12222222222220580079},
		{12312.0, 7095.37222222222226264421},
		{0.0000043, 255.37222461111107918441},
		{1234.0, 940.92777777777780556789},
		{659846935.0, 366581885.92777776718139648438},
		{1054004.0, 585813.15000000002328306437},
		{-100.432423, 199.57643166666665024422},
	}

	for _, test := range tests {
		result := FahrenheitToKelvin(test.input)
		if !almostEqualK(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	tests := []struct {
		input    Kelvin
		expected Fahrenheit
	}{
		{0.0, -459.66999999999995907274},
		{100.0, -279.66999999999995907274},
		{273.15, 32.0},
		{12312.0, 21701.93000000000029103830},
		{0.0000043, -459.66999225999995815073},
		{1234.0, 1761.52999999999997271516},
		{659846935.0, 1187724023.33000016212463378906},
		{1054004.0, 1896747.53000000026077032089},
		{-100.432423, -640.44836139999995339167},
	}

	for _, test := range tests {
		result := KelvinToFahrenheit(test.input)
		if !almostEqualF(result, test.expected) {
			t.Errorf("expected %.20f, but got %.20f", test.expected, result)
		}
	}
}
