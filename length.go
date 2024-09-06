package main

var LengthUnits = map[string]float64{
	"kilometer":  1000,
	"hectometer": 100,
	"decameter":  10,
	"meter":      1,
	"decimeter":  0.1,
	"centimeter": 0.01,
	"millimeter": 0.001,
	"mile":       1609.34,
	"yard":       0.9144,
	"foot":       0.3048,
	"inch":       0.0254,
}

func ConvertLength(value float64, fromUnit string, toUnit string) float64 {
	result := 0

	if fromUnit == "" || toUnit == "" {
		return value
	}

	return float64(result)
}
