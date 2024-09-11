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

func GigatonneToMegatonne(gigatonne float64) float64 {
	return gigatonne * 1000
}

func GigatonneToTonne(gigatonne float64) float64 {
	return gigatonne * 1000000000
}

func GigatonneToKilogram(gigatonne float64) float64 {
	return gigatonne * 1000000000000
}

func GigatonneToPound(gigatonne float64) float64 {
	return gigatonne * 2204622621848.78
}

func GigatonneToGram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000
}

func GigatonneToMilligram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000000
}

func GigatonneToMicrogram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000000000
}

func GigatonneToNanogram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000000000000
}

func GigatonneToPicogram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000000000000000
}

func MegatonneToGigatonne(megatonne float64) float64 {
	return megatonne * 0.001
}

func MegatonneToTonne(megatonne float64) float64 {
	return megatonne * 1000000
}

func MegatonneToKilogram(megatonne float64) float64 {
	return megatonne * 1000000000
}

func MegatonneToPound(megatonne float64) float64 {
	return megatonne * 2204622621.85
}

func MegatonneToGram(megatonne float64) float64 {
	return megatonne * 1000000000000
}

func MegatonneToMilligram(megatonne float64) float64 {
	return megatonne * 1000000000000000
}

func MegatonneToMicrogram(megatonne float64) float64 {
	return megatonne * 1000000000000000000
}

func MegatonneToNanogram(megatonne float64) float64 {
	return megatonne * 1000000000000000000000
}

func MegatonneToPicogram(megatonne float64) float64 {
	return megatonne * 1000000000000000000000000
}