package weight

import (
	"fmt"
)

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
func ConvertWeight(value float64, fromUnit string, toUnit string) (float64, error) {
	fromUnitValue, fromExists := weightUnits[fromUnit]
	if !fromExists {
		return 0, fmt.Errorf("invalid from unit: %s", fromUnit)
	}

	toUnitValue, toExists := weightUnits[toUnit]
	if !toExists {
		return 0, fmt.Errorf("invalid to unit: %s", toUnit)
	}

	fromValueInGram := value * fromUnitValue
	result := fromValueInGram / toUnitValue

	return result, nil
}

// Convert Gigatonne to Megatonne
func GigatonneToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "megatonne")
}

// Convert Gigatonne to Tonne
func GigatonneToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "tonne")
}

// Convert Gigatonne to Kilogram
func GigatonneToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "kilogram")
}

// Convert Gigatonne to Pound
func GigatonneToPound(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "pound")
}

// Convert Gigatonne to Gram
func GigatonneToGram(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "gram")
}

// Convert Gigatonne to Milligram
func GigatonneToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "milligram")
}

// Convert Gigatonne to Microgram
func GigatonneToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "microgram")
}

// Convert Gigatonne to Nanogram
func GigatonneToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "nanogram")
}

// Convert Gigatonne to Picogram
func GigatonneToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "gigatonne", "picogram")
}

// Convert Megatonne to Gigatonne
func MegatonneToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "gigatonne")
}

// Convert Megatonne to Tonne
func MegatonneToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "tonne")
}

// Convert Megatonne to Kilogram
func MegatonneToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "kilogram")
}

// Convert Megatonne to Pound
func MegatonneToPound(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "pound")
}

// Convert Megatonne to Gram
func MegatonneToGram(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "gram")
}

// Convert Megatonne to Milligram
func MegatonneToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "milligram")
}

// Convert Megatonne to Microgram
func MegatonneToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "microgram")
}

// Convert Megatonne to Nanogram
func MegatonneToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "nanogram")
}

// Convert Megatonne to Picogram
func MegatonneToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "megatonne", "picogram")
}

// Convert Tonne to Gigatonne
func TonneToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "gigatonne")
}

// Convert Tonne to Megatonne
func TonneToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "megatonne")
}

// Convert Tonne to Kilogram
func TonneToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "kilogram")
}

// Convert Tonne to Pound
func TonneToPound(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "pound")
}

// Convert Tonne to Gram
func TonneToGram(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "gram")
}

// Convert Tonne to Milligram
func TonneToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "milligram")
}

// Convert Tonne to Microgram
func TonneToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "microgram")
}

// Convert Tonne to Nanogram
func TonneToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "nanogram")
}

// Convert Tonne to Picogram
func TonneToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "tonne", "picogram")
}

// Convert Kilogram to Gigatonne
func KilogramToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "gigatonne")
}

// Convert Kilogram to Megatonne
func KilogramToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "megatonne")
}

// Convert Kilogram to Tonne
func KilogramToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "tonne")
}

// Convert Kilogram to Pound
func KilogramToPound(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "pound")
}

// Convert Kilogram to Gram
func KilogramToGram(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "gram")
}

// Convert Kilogram to Milligram
func KilogramToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "milligram")
}

// Convert Kilogram to Microgram
func KilogramToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "microgram")
}

// Convert Kilogram to Nanogram
func KilogramToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "nanogram")
}

// Convert Kilogram to Picogram
func KilogramToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "kilogram", "picogram")
}

// Convert Pound to Gigatonne
func PoundToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "gigatonne")
}

// Convert Pound to Megatonne
func PoundToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "megatonne")
}

// Convert Pound to Tonne
func PoundToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "tonne")
}

// Convert Pound to Kilogram
func PoundToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "kilogram")
}

// Convert Pound to Gram
func PoundToGram(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "gram")
}

// Convert Pound to Milligram
func PoundToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "milligram")
}

// Convert Pound to Microgram
func PoundToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "microgram")
}

// Convert Pound to Nanogram
func PoundToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "nanogram")
}

// Convert Pound to Picogram
func PoundToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "pound", "picogram")
}

// Convert Gram to Gigatonne
func GramToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "gigatonne")
}

// Convert Gram to Megatonne
func GramToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "megatonne")
}

// Convert Gram to Tonne
func GramToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "tonne")
}

// Convert Gram to Kilogram
func GramToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "kilogram")
}

// Convert Gram to Pound
func GramToPound(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "pound")
}

// Convert Gram to Milligram
func GramToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "milligram")
}

// Convert Gram to Microgram
func GramToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "microgram")
}

// Convert Gram to Nanogram
func GramToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "nanogram")
}

// Convert Gram to Picogram
func GramToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "gram", "picogram")
}

// Convert Milligram to Gigatonne
func MilligramToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "gigatonne")
}

// Convert Milligram to Megatonne
func MilligramToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "megatonne")
}

// Convert Milligram to Tonne
func MilligramToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "tonne")
}

// Convert Milligram to Kilogram
func MilligramToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "kilogram")
}

// Convert Milligram to Pound
func MilligramToPound(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "pound")
}

// Convert Milligram to Gram
func MilligramToGram(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "gram")
}

// Convert Milligram to Microgram
func MilligramToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "microgram")
}

// Convert Milligram to Nanogram
func MilligramToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "nanogram")
}

// Convert Milligram to Picogram
func MilligramToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "milligram", "picogram")
}

// Convert Microgram to Gigatonne
func MicrogramToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "gigatonne")
}

// Convert Microgram to Megatonne
func MicrogramToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "megatonne")
}

// Convert Microgram to Tonne
func MicrogramToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "tonne")
}

// Convert Microgram to Kilogram
func MicrogramToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "kilogram")
}

// Convert Microgram to Pound
func MicrogramToPound(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "pound")
}

// Convert Microgram to Gram
func MicrogramToGram(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "gram")
}

// Convert Microgram to Milligram
func MicrogramToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "milligram")
}

// Convert Microgram to Nanogram
func MicrogramToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "nanogram")
}

// Convert Microgram to Picogram
func MicrogramToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "microgram", "picogram")
}

// Convert Nanogram to Gigatonne
func NanogramToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "gigatonne")
}

// Convert Nanogram to Megatonne
func NanogramToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "megatonne")
}

// Convert Nanogram to Tonne
func NanogramToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "tonne")
}

// Convert Nanogram to Kilogram
func NanogramToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "kilogram")
}

// Convert Nanogram to Pound
func NanogramToPound(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "pound")
}

// Convert Nanogram to Gram
func NanogramToGram(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "gram")
}

// Convert Nanogram to Milligram
func NanogramToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "milligram")
}

// Convert Nanogram to Microgram
func NanogramToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "microgram")
}

// Convert Nanogram to Picogram
func NanogramToPicogram(value float64) (float64, error) {
	return ConvertWeight(value, "nanogram", "picogram")
}

// Convert Picogram to Gigatonne
func PicogramToGigatonne(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "gigatonne")
}

// Convert Picogram to Megatonne
func PicogramToMegatonne(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "megatonne")
}

// Convert Picogram to Tonne
func PicogramToTonne(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "tonne")
}

// Convert Picogram to Kilogram
func PicogramToKilogram(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "kilogram")
}

// Convert Picogram to Pound
func PicogramToPound(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "pound")
}

// Convert Picogram to Gram
func PicogramToGram(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "gram")
}

// Convert Picogram to Milligram
func PicogramToMilligram(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "milligram")
}

// Convert Picogram to Microgram
func PicogramToMicrogram(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "microgram")
}

// Convert Picogram to Nanogram
func PicogramToNanogram(value float64) (float64, error) {
	return ConvertWeight(value, "picogram", "nanogram")
}
