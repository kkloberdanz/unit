package unit

type Meter float64
type Angstrom float64

type Inch float64
type Foot float64
type Yard float64
type Mile float64

func MeterToAngstrom(meters Meter) Angstrom {
	return Angstrom(1e-10 * meters)
}

func MeterToInch(meters Meter) Inch {
	return Inch(meters * 39.37008)
}

func MeterToFoot(meters Meter) Foot {
	return Foot(meters * 3.28084)
}

func MeterToYard(meters Meter) Yard {
	return Yard(meters * 1.093613)
}

func MeterToMile(meters Meter) Mile {
	return Mile(meters * 0.0006213712)
}
