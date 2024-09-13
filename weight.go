package convertunits

import "fmt"

var weightUnits = map[string]float64{
	"gigatonne": 1e15,
	"megatonne": 1e12,
	"tonne":     1e6,
	"kilogram":  1000,
	"pound":     453.592,
	"gram":      1,
	"milligram": 0.001,
	"microgram": 1e-6,
	"nanogram":  1e-9,
	"picogram":  1e-12,
}

// ConvertWeight converts the given weight from one unit to another
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

// Convert Gigatonne to Megatonne
func GigatonneToMegatonne(gigatonne float64) float64 {
	return gigatonne * 1000
}

// Convert Gigatonne to Tonne
func GigatonneToTonne(gigatonne float64) float64 {
	return gigatonne * 1000000000
}

// Convert Gigatonne to Kilogram
func GigatonneToKilogram(gigatonne float64) float64 {
	return gigatonne * 1000000000000
}

// Convert Gigatonne to Pound
func GigatonneToPound(gigatonne float64) float64 {
	return gigatonne * 2204622621848.78
}

// Convert Gigatonne to Gram
func GigatonneToGram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000
}

// Convert Gigatonne to Milligram
func GigatonneToMilligram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000000
}

// Convert Gigatonne to Microgram
func GigatonneToMicrogram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000000000
}

// Convert Gigatonne to Nanogram
func GigatonneToNanogram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000000000000
}

// Convert Gigatonne to Picogram
func GigatonneToPicogram(gigatonne float64) float64 {
	return gigatonne * 1000000000000000000000000000
}

// Convert Megatonne to Gigatonne
func MegatonneToGigatonne(megatonne float64) float64 {
	return megatonne * 0.001
}

// Convert Megatonne to Tonne
func MegatonneToTonne(megatonne float64) float64 {
	return megatonne * 1000000
}

// Convert Megatonne to Kilogram
func MegatonneToKilogram(megatonne float64) float64 {
	return megatonne * 1000000000
}

// Convert Megatonne to Pound
func MegatonneToPound(megatonne float64) float64 {
	return megatonne * 2204622621.85
}

// Convert Megatonne to Gram
func MegatonneToGram(megatonne float64) float64 {
	return megatonne * 1000000000000
}

// Convert Megatonne to Milligram
func MegatonneToMilligram(megatonne float64) float64 {
	return megatonne * 1000000000000000
}

// Convert Megatonne to Microgram
func MegatonneToMicrogram(megatonne float64) float64 {
	return megatonne * 1000000000000000000
}

// Convert Megatonne to Nanogram
func MegatonneToNanogram(megatonne float64) float64 {
	return megatonne * 1000000000000000000000
}

// Convert Megatonne to Picogram
func MegatonneToPicogram(megatonne float64) float64 {
	return megatonne * 1000000000000000000000000
}

// Convert Tonne to Gigatonne
func TonneToGigatonne(tonne float64) float64 {
	return tonne * 0.000000001
}

// Convert Tonne to Megatonne
func TonneToMegatonne(tonne float64) float64 {
	return tonne * 0.000001
}

// Convert Tonne to Kilogram
func TonneToKilogram(tonne float64) float64 {
	return tonne * 1000
}

// Convert Tonne to Pound
func TonneToPound(tonne float64) float64 {
	return tonne * 2204.62262185
}

// Convert Tonne to Gram
func TonneToGram(tonne float64) float64 {
	return tonne * 1000000
}

// Convert Tonne to Milligram
func TonneToMilligram(tonne float64) float64 {
	return tonne * 1000000000
}

// Convert Tonne to Microgram
func TonneToMicrogram(tonne float64) float64 {
	return tonne * 1000000000000
}

// Convert Tonne to Nanogram
func TonneToNanogram(tonne float64) float64 {
	return tonne * 1000000000000000
}

// Convert Tonne to Picogram
func TonneToPicogram(tonne float64) float64 {
	return tonne * 1000000000000000000
}

// Convert Kilogram to Gigatonne
func KilogramToGigatonne(kilogram float64) float64 {
	return kilogram * 0.000000000001
}

// Convert Kilogram to Megatonne
func KilogramToMegatonne(kilogram float64) float64 {
	return kilogram * 0.000000001
}

// Convert Kilogram to Tonne
func KilogramToTonne(kilogram float64) float64 {
	return kilogram * 0.001
}

// Convert Kilogram to Pound
func KilogramToPound(kilogram float64) float64 {
	return kilogram * 2.20462262185
}

// Convert Kilogram to Gram
func KilogramToGram(kilogram float64) float64 {
	return kilogram * 1000
}

// Convert Kilogram to Milligram
func KilogramToMilligram(kilogram float64) float64 {
	return kilogram * 1000000
}

// Convert Kilogram to Microgram
func KilogramToMicrogram(kilogram float64) float64 {
	return kilogram * 1000000000
}

// Convert Kilogram to Nanogram
func KilogramToNanogram(kilogram float64) float64 {
	return kilogram * 1000000000000
}

// Convert Kilogram to Picogram
func KilogramToPicogram(kilogram float64) float64 {
	return kilogram * 1000000000000000
}

// Convert Pound to Gigatonne
func PoundToGigatonne(pound float64) float64 {
	return pound * 0.00000000000045359237
}

// Convert Pound to Megatonne
func PoundToMegatonne(pound float64) float64 {
	return pound * 0.00000000045359237
}

// Convert Pound to Tonne
func PoundToTonne(pound float64) float64 {
	return pound * 0.00045359237
}

// Convert Pound to Kilogram
func PoundToKilogram(pound float64) float64 {
	return pound * 0.45359237
}

// Convert Pound to Gram
func PoundToGram(pound float64) float64 {
	return pound * 453.59237
}

// Convert Pound to Milligram
func PoundToMilligram(pound float64) float64 {
	return pound * 453592.37
}

// Convert Pound to Microgram
func PoundToMicrogram(pound float64) float64 {
	return pound * 453592370
}

// Convert Pound to Nanogram
func PoundToNanogram(pound float64) float64 {
	return pound * 453592370000
}

// Convert Pound to Picogram
func PoundToPicogram(pound float64) float64 {
	return pound * 453592370000000
}

// Convert Gram to Gigatonne
func GramToGigatonne(gram float64) float64 {
	return gram * 0.000000000000001
}

// Convert Gram to Megatonne
func GramToMegatonne(gram float64) float64 {
	return gram * 0.000000000001
}

// Convert Gram to Tonne
func GramToTonne(gram float64) float64 {
	return gram * 0.000001
}

// Convert Gram to Kilogram
func GramToKilogram(gram float64) float64 {
	return gram * 0.001
}

// Convert Gram to Pound
func GramToPound(gram float64) float64 {
	return gram * 0.00220462262185
}

// Convert Gram to Milligram
func GramToMilligram(gram float64) float64 {
	return gram * 1000
}

// Convert Gram to Microgram
func GramToMicrogram(gram float64) float64 {
	return gram * 1000000
}

// Convert Gram to Nanogram
func GramToNanogram(gram float64) float64 {
	return gram * 1000000000
}

// Convert Gram to Picogram
func GramToPicogram(gram float64) float64 {
	return gram * 1000000000000
}

// Convert Milligram to Gigatonne
func MilligramToGigatonne(milligram float64) float64 {
	return milligram * 0.000000000000000001
}

// Convert Milligram to Megatonne
func MilligramToMegatonne(milligram float64) float64 {
	return milligram * 0.000000000000001
}

// Convert Milligram to Tonne
func MilligramToTonne(milligram float64) float64 {
	return milligram * 0.000000001
}

// Convert Milligram to Kilogram
func MilligramToKilogram(milligram float64) float64 {
	return milligram * 0.000001
}

// Convert Milligram to Pound
func MilligramToPound(milligram float64) float64 {
	return milligram * 0.00000220462262185
}

// Convert Milligram to Gram
func MilligramToGram(milligram float64) float64 {
	return milligram * 0.001
}

// Convert Milligram to Microgram
func MilligramToMicrogram(milligram float64) float64 {
	return milligram * 1000
}

// Convert Milligram to Nanogram
func MilligramToNanogram(milligram float64) float64 {
	return milligram * 1000000
}

// Convert Milligram to Picogram
func MilligramToPicogram(milligram float64) float64 {
	return milligram * 1000000000
}

// Convert Microgram to Gigatonne
func MicrogramToGigatonne(microgram float64) float64 {
	return microgram * 0.000000000000000000001
}

// Convert Microgram to Megatonne
func MicrogramToMegatonne(microgram float64) float64 {
	return microgram * 0.000000000000000001
}

// Convert Microgram to Tonne
func MicrogramToTonne(microgram float64) float64 {
	return microgram * 0.000000000001
}

// Convert Microgram to Kilogram
func MicrogramToKilogram(microgram float64) float64 {
	return microgram * 0.000000001
}

// Convert Microgram to Pound
func MicrogramToPound(microgram float64) float64 {
	return microgram * 0.00000000220462262185
}

// Convert Microgram to Gram
func MicrogramToGram(microgram float64) float64 {
	return microgram * 0.000001
}

// Convert Microgram to Milligram
func MicrogramToMilligram(microgram float64) float64 {
	return microgram * 0.001
}

// Convert Microgram to Nanogram
func MicrogramToNanogram(microgram float64) float64 {
	return microgram * 1000
}

// Convert Microgram to Picogram
func MicrogramToPicogram(microgram float64) float64 {
	return microgram * 1000000
}

// Convert Nanogram to Gigatonne
func NanogramToGigatonne(nanogram float64) float64 {
	return nanogram * 0.000000000000000000000001
}

// Convert Nanogram to Megatonne
func NanogramToMegatonne(nanogram float64) float64 {
	return nanogram * 0.000000000000000000001
}

// Convert Nanogram to Tonne
func NanogramToTonne(nanogram float64) float64 {
	return nanogram * 0.000000000000001
}

// Convert Nanogram to Kilogram
func NanogramToKilogram(nanogram float64) float64 {
	return nanogram * 0.000000000001
}

// Convert Nanogram to Pound
func NanogramToPound(nanogram float64) float64 {
	return nanogram * 0.00000000000220462262185
}

// Convert Nanogram to Gram
func NanogramToGram(nanogram float64) float64 {
	return nanogram * 0.000000001
}

// Convert Nanogram to Milligram
func NanogramToMilligram(nanogram float64) float64 {
	return nanogram * 0.000001
}

// Convert Nanogram to Microgram
func NanogramToMicrogram(nanogram float64) float64 {
	return nanogram * 0.001
}

// Convert Nanogram to Picogram
func NanogramToPicogram(nanogram float64) float64 {
	return nanogram * 1000
}

// Convert Picogram to Gigatonne
func PicogramToGigatonne(picogram float64) float64 {
	return picogram * 0.000000000000000000000000001
}

// Convert Picogram to Megatonne
func PicogramToMegatonne(picogram float64) float64 {
	return picogram * 0.000000000000000000000001
}

// Convert Picogram to Tonne
func PicogramToTonne(picogram float64) float64 {
	return picogram * 0.000000000000000001
}

// Convert Picogram to Kilogram
func PicogramToKilogram(picogram float64) float64 {
	return picogram * 0.000000000000001
}

// Convert Picogram to Pound
func PicogramToPound(picogram float64) float64 {
	return picogram * 0.00000000000000000220462262185
}

// Convert Picogram to Gram
func PicogramToGram(picogram float64) float64 {
	return picogram * 0.000000000001
}

// Convert Picogram to Milligram
func PicogramToMilligram(picogram float64) float64 {
	return picogram * 0.000000001
}

// Convert Picogram to Microgram
func PicogramToMicrogram(picogram float64) float64 {
	return picogram * 0.000001
}

// Convert Picogram to Nanogram
func PicogramToNanogram(picogram float64) float64 {
	return picogram * 0.001
}
