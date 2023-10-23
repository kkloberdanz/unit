package temperature

type Celcius float64
type Farenheit float64

func CelciusToFarenheit(temp Celcius) Farenheit {
	return Farenheit(temp*Celcius(9.0/5.0) + Celcius(32.0))
}

func FarenheitToCelcius(temp Farenheit) Celcius {
	return Celcius((temp - Farenheit(32.0)) * Farenheit(5.0/9.0))
}
