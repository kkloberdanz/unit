package temperature

type Kelvin float64
type Celcius float64
type Fahrenheit float64

func KelvinToCelcius(temp Kelvin) Celcius {
	return Celcius(temp - 273.15)
}

func KelvinToFahrenheit(temp Kelvin) Fahrenheit {
	return Fahrenheit((temp-273.15)*(9.0/5.0) + 32.0)
}

func CelciusToKelvin(temp Celcius) Kelvin {
	return Kelvin(temp + 273.15)
}

func CelciusToFahrenheit(temp Celcius) Fahrenheit {
	return Fahrenheit(temp*(9.0/5.0) + 32.0)
}

func FahrenheitToCelcius(temp Fahrenheit) Celcius {
	return Celcius((temp - 32.0) * (5.0 / 9.0))
}

func FahrenheitToKelvin(temp Fahrenheit) Kelvin {
	return Kelvin((temp-32.0)*(5.0/9.0) + 273.15)
}
