package temperature

type Kelvin float64
type Celcius float64
type Farenheit float64

func KelvinToCelcius(temp Kelvin) Celcius {
	return Celcius(temp - 273.15)
}

func CelciusToFarenheit(temp Celcius) Farenheit {
	return Farenheit(temp*(9.0/5.0) + 32.0)
}

func FarenheitToCelcius(temp Farenheit) Celcius {
	return Celcius((temp - 32.0) * (5.0 / 9.0))
}
