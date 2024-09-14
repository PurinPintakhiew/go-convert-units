package temperature

import "fmt"

type tempFunc func(value float64) (float64, error)

var tempConversions = map[string]map[string]tempFunc{
	"celsius": {
		"fahrenheit": CelsiusToFahrenheit,
		"kelvin":     CelsiusToKelvin,
		"rankine":    CelsiusToRankine,
		"reaumur":    CelsiusToReaumur,
	},
	"fahrenheit": {
		"celsius": FahrenheitToCelsius,
		"kelvin":  FahrenheitToKelvin,
		"rankine": FahrenheitToRankine,
		"reaumur": FahrenheitToReaumur,
	},
	"kelvin": {
		"celsius":    KelvinToCelsius,
		"fahrenheit": KelvinToFahrenheit,
		"rankine":    KelvinToRankine,
		"reaumur":    KelvinToReaumur,
	},
	"rankine": {
		"celsius":    RankineToCelsius,
		"fahrenheit": RankineToFahrenheit,
		"kelvin":     RankineToKelvin,
		"reaumur":    RankineToReaumur,
	},
	"reaumur": {
		"celsius":    ReaumurToCelsius,
		"fahrenheit": ReaumurToFahrenheit,
		"kelvin":     ReaumurToKelvin,
		"rankine":    ReaumurToRankine,
	},
}

// ConvertTemperature converts the given Temperature from one unit to another
func ConvertTemperature(value float64, fromUnit string, toUnit string) (float64, error) {
	fromConversions, fromExists := tempConversions[fromUnit]
	if !fromExists {
		return 0, fmt.Errorf("invalid from unit: %s", fromUnit)
	}

	toConversion, toExists := fromConversions[toUnit]
	if !toExists {
		return 0, fmt.Errorf("invalid to unit: %s", toUnit)
	}

	result, err := toConversion(value)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// Convert Celsius to Fahrenheit
func CelsiusToFahrenheit(celsius float64) (float64, error) {
	return (celsius * 9 / 5) + 32, nil
}

// Convert Celsius to Kelvin
func CelsiusToKelvin(celsius float64) (float64, error) {
	return celsius + 273.15, nil
}

// Convert Celsius to Rankine
func CelsiusToRankine(celsius float64) (float64, error) {
	return (celsius + 273.15) * 9 / 5, nil
}

// Convert Celsius to Reaumur
func CelsiusToReaumur(celsius float64) (float64, error) {
	return celsius * 4 / 5, nil
}

// Convert Fahrenheit to Celsius
func FahrenheitToCelsius(fahrenheit float64) (float64, error) {
	return (fahrenheit - 32) * 5 / 9, nil
}

// Convert Fahrenheit to Kelvin
func FahrenheitToKelvin(fahrenheit float64) (float64, error) {
	return (fahrenheit-32)*5/9 + 273.15, nil
}

// Convert Fahrenheit to Rankine
func FahrenheitToRankine(fahrenheit float64) (float64, error) {
	return fahrenheit + 459.67, nil
}

// Convert Fahrenheit to Reaumur
func FahrenheitToReaumur(fahrenheit float64) (float64, error) {
	return (fahrenheit - 32) * 4 / 9, nil
}

// Convert Kelvin to Celsius
func KelvinToCelsius(kelvin float64) (float64, error) {
	return kelvin - 273.15, nil
}

// Convert Kelvin to Fahrenheit
func KelvinToFahrenheit(kelvin float64) (float64, error) {
	return (kelvin-273.15)*9/5 + 32, nil
}

// Convert Kelvin to Rankine
func KelvinToRankine(kelvin float64) (float64, error) {
	return kelvin * 9 / 5, nil
}

// Convert Kelvin to Reaumur
func KelvinToReaumur(kelvin float64) (float64, error) {
	return (kelvin - 273.15) * 4 / 5, nil
}

// Convert Rankine to Celsius
func RankineToCelsius(rankine float64) (float64, error) {
	return (rankine - 491.67) * 5 / 9, nil
}

// Convert Rankine to Fahrenheit
func RankineToFahrenheit(rankine float64) (float64, error) {
	return rankine - 459.67, nil
}

// Convert Rankine to Kelvin
func RankineToKelvin(rankine float64) (float64, error) {
	return rankine * 5 / 9, nil
}

// Convert Rankine to Reaumur
func RankineToReaumur(rankine float64) (float64, error) {
	return (rankine - 491.67) * 4 / 9, nil
}

// Convert Reaumur to Celsius
func ReaumurToCelsius(reaumur float64) (float64, error) {
	return reaumur * 5 / 4, nil
}

// Convert Reaumur to Fahrenheit
func ReaumurToFahrenheit(reaumur float64) (float64, error) {
	return (reaumur * 9 / 4) + 32, nil
}

// Convert Reaumur to Kelvin
func ReaumurToKelvin(reaumur float64) (float64, error) {
	return (reaumur * 5 / 4) + 273.15, nil
}

// Convert Reaumur to Rankine
func ReaumurToRankine(reaumur float64) (float64, error) {
	return (reaumur * 9 / 4) + 491.67, nil
}
