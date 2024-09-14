package temperature

import "fmt"

// ConvertTemperature converts the given Temperature from one unit to another
func ConvertTemperature(value float64, fromUnit string, toUnit string) float64 {
	switch fromUnit {
	case "celsius":
		return convertFromCelsius(value, toUnit)
	case "fahrenheit":
		return convertFromFahrenheit(value, toUnit)
	case "kelvin":
		return convertFromKelvin(value, toUnit)
	case "rankine":
		return convertFromRankine(value, toUnit)
	case "reaumur":
		return convertFromReaumur(value, toUnit)
	default:
		fmt.Printf("Invalid units: %s\n", fromUnit)
		return 0.0
	}
}

func convertFromCelsius(value float64, toUnit string) float64 {
	switch toUnit {
	case "fahrenheit":
		return CelsiusToFahrenheit(value)
	case "kelvin":
		return CelsiusToKelvin(value)
	case "rankine":
		return CelsiusToRankine(value)
	case "reaumur":
		return CelsiusToReaumur(value)
	case "celsius":
		return value
	default:
		fmt.Printf("Invalid units: %s\n", toUnit)
		return 0.0
	}
}

func convertFromFahrenheit(value float64, toUnit string) float64 {
	switch toUnit {
	case "celsius":
		return FahrenheitToCelsius(value)
	case "kelvin":
		return FahrenheitToKelvin(value)
	case "rankine":
		return FahrenheitToRankine(value)
	case "reaumur":
		return FahrenheitToReaumur(value)
	case "fahrenheit":
		return value
	default:
		fmt.Printf("Invalid units: %s\n", toUnit)
		return 0.0
	}
}

func convertFromKelvin(value float64, toUnit string) float64 {
	switch toUnit {
	case "celsius":
		return KelvinToCelsius(value)
	case "fahrenheit":
		return KelvinToFahrenheit(value)
	case "rankine":
		return KelvinToRankine(value)
	case "reaumur":
		return KelvinToReaumur(value)
	case "kelvin":
		return value
	default:
		fmt.Printf("Invalid units: %s\n", toUnit)
		return 0.0
	}
}

func convertFromRankine(value float64, toUnit string) float64 {
	switch toUnit {
	case "celsius":
		return RankineToCelsius(value)
	case "fahrenheit":
		return RankineToFahrenheit(value)
	case "kelvin":
		return RankineToKelvin(value)
	case "reaumur":
		return RankineToReaumur(value)
	case "rankine":
		return value
	default:
		fmt.Printf("Invalid units: %s\n", toUnit)
		return 0.0
	}
}

func convertFromReaumur(value float64, toUnit string) float64 {
	switch toUnit {
	case "celsius":
		return ReaumurToCelsius(value)
	case "fahrenheit":
		return ReaumurToFahrenheit(value)
	case "kelvin":
		return ReaumurToKelvin(value)
	case "rankine":
		return ReaumurToRankine(value)
	case "reaumur":
		return value
	default:
		fmt.Printf("Invalid units: %s\n", toUnit)
		return 0.0
	}
}

// Convert Celsius to Fahrenheit
func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// Convert Celsius to Kelvin
func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// Convert Celsius to Rankine
func CelsiusToRankine(celsius float64) float64 {
	return (celsius + 273.15) * 9 / 5
}

// Convert Celsius to Reaumur
func CelsiusToReaumur(celsius float64) float64 {
	return celsius * 4 / 5
}

// Convert Fahrenheit to Celsius
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// Convert Fahrenheit to Kelvin
func FahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273.15
}

// Convert Fahrenheit to Rankine
func FahrenheitToRankine(fahrenheit float64) float64 {
	return fahrenheit + 459.67
}

// Convert Fahrenheit to Reaumur
func FahrenheitToReaumur(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 4 / 9
}

// Convert Kelvin to Celsius
func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

// Convert Kelvin to Fahrenheit
func KelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

// Convert Kelvin to Rankine
func KelvinToRankine(kelvin float64) float64 {
	return kelvin * 9 / 5
}

// Convert Kelvin to Reaumur
func KelvinToReaumur(kelvin float64) float64 {
	return (kelvin - 273.15) * 4 / 5
}

// Convert Rankine to Celsius
func RankineToCelsius(rankine float64) float64 {
	return (rankine - 491.67) * 5 / 9
}

// Convert Rankine to Fahrenheit
func RankineToFahrenheit(rankine float64) float64 {
	return rankine - 459.67
}

// Convert Rankine to Kelvin
func RankineToKelvin(rankine float64) float64 {
	return rankine * 5 / 9
}

// Convert Rankine to Reaumur
func RankineToReaumur(rankine float64) float64 {
	return (rankine - 491.67) * 4 / 9
}

// Convert Reaumur to Celsius
func ReaumurToCelsius(reaumur float64) float64 {
	return reaumur * 5 / 4
}

// Convert Reaumur to Fahrenheit
func ReaumurToFahrenheit(reaumur float64) float64 {
	return (reaumur * 9 / 4) + 32
}

// Convert Reaumur to Kelvin
func ReaumurToKelvin(reaumur float64) float64 {
	return (reaumur * 5 / 4) + 273.15
}

// Convert Reaumur to Rankine
func ReaumurToRankine(reaumur float64) float64 {
	return (reaumur * 9 / 4) + 491.67
}
