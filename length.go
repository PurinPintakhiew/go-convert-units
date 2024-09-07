package main

import "fmt"

var lengthUnits = map[string]float64{
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
	fromUnitValue, fromExists := lengthUnits[fromUnit]
	toUnitValue, toExists := lengthUnits[toUnit]

	if !fromExists || !toExists {
		fmt.Printf("Invalid units: %s or %s\n", fromUnit, toUnit)
		return 0
	}

	// --- calc
	fromValueInMeter := value * fromUnitValue
	result := fromValueInMeter / toUnitValue

	return result
}
