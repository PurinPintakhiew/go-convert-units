package convertunits

import "fmt"

var weightUnits = map[string]float64{
	"gigatonne": 1000000000000000,
	"megatonne": 1000000000000,
	"tonne":     1000000,
	"kilogram":  1000,
	"pound":     453.592,
	"gram":      1,
	"milligram": 0.001,
	"microgram": 0.000001,
	"nanogram":  0.000000001,
	"picogram":  0.000000000001,
}

func ConvertWeight(value float64, fromUnit string, toUnit string) float64 {
	fromUnitValue, fromExists := weightUnits[fromUnit]
	toUnitValue, toExists := weightUnits[toUnit]

	if !fromExists || !toExists {
		fmt.Printf("Invalid units: %s or %s\n", fromUnit, toUnit)
		return 0
	}

	// --- calc
	fromValueInGram := value * fromUnitValue
	result := fromValueInGram / toUnitValue

	return result
}
