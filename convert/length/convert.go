package length

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

// ConvertLength converts the given length from one unit to another
func ConvertLength(value float64, fromUnit string, toUnit string) (float64, error) {
	fromUnitValue, fromExists := lengthUnits[fromUnit]
	if !fromExists {
		return 0, fmt.Errorf("invalid from unit: %s", fromUnit)
	}

	toUnitValue, toExists := lengthUnits[toUnit]
	if !toExists {
		return 0, fmt.Errorf("invalid to unit: %s", toUnit)
	}

	fromValueInMeter := value * fromUnitValue
	result := fromValueInMeter / toUnitValue

	return result, nil
}

// Convert Kilometer to Hectometer
func KilometerToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "hectometer")
}

// Convert Kilometer to Decameter
func KilometerToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "decameter")
}

// Convert Kilometer to Meter
func KilometerToMeter(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "meter")
}

// Convert Kilometer to Meter
func KilometerToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "decimeter")
}

// Convert Kilometer to Centimeter
func KilometerToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "centimeter")
}

// Convert Kilometer to Millimeter
func KilometerToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "millimeter")
}

// Convert Kilometer to Mile
func KilometerToMile(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "mile")
}

// Convert Kilometer to Yard
func KilometerToYard(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "yard")
}

// Convert Kilometer to Foot
func KilometerToFoot(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "foot")
}

// Convert Kilometer to Inch
func KilometerToInch(value float64) (float64, error) {
	return ConvertLength(value, "kilometer", "inch")
}

// Convert Hectometer to kilometer
func HectometerTokilometer(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "kilometer")
}

// Convert Hectometer to Decameter
func HectometerToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "decameter")
}

// Convert Hectometer to Meter
func HectometerToMeter(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "meter")
}

// Convert Hectometer to Decimeter
func HectometerToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "decimeter")
}

// Convert Hectometer to Centimeter
func HectometerToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "centimeter")
}

// Convert Hectometer to Millimeter
func HectometerToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "millimeter")
}

// Convert Hectometer to Mile
func HectometerToMile(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "mile")
}

// Convert Hectometer to Yard
func HectometerToYard(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "yard")
}

// Convert Hectometer to Foot
func HectometerToFoot(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "foot")
}

// Convert Hectometer to Inch
func HectometerToInch(value float64) (float64, error) {
	return ConvertLength(value, "hectometer", "inch")
}

// Convert Decameter to Kilometer
func DecameterToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "kilometer")
}

// Convert Decameter to Hectometer
func DecameterToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "hectometer")
}

// Convert Decameter to Meter
func DecameterToMeter(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "meter")
}

// Convert Decameter to Decimeter
func DecameterToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "decimeter")
}

// Convert Decameter to Centimeter
func DecameterToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "centimeter")
}

// Convert Decameter to Millimeter
func DecameterToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "millimeter")
}

// Convert Decameter to Mile
func DecameterToMile(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "mile")
}

// Convert Decameter to Yard
func DecameterToYard(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "yard")
}

// Convert Decameter to Foot
func DecameterToFoot(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "foot")
}

// Convert Decameter to Inch
func DecameterToInch(value float64) (float64, error) {
	return ConvertLength(value, "decameter", "inch")
}

// Convert Meter to Kilometer
func MeterToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "meter", "kilometer")
}

// Convert Meter to Hectometer
func MeterToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "meter", "hectometer")
}

// Convert Meter to Decameter
func MeterToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "meter", "decameter")
}

// Convert Meter to Decimeter
func MeterToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "meter", "decimeter")
}

// Convert Meter to Centimeter
func MeterToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "meter", "centimeter")
}

// Convert Meter to Millimeter
func MeterToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "meter", "millimeter")
}

// Convert Meter to Mile
func MeterToMile(value float64) (float64, error) {
	return ConvertLength(value, "meter", "mile")
}

// Convert Meter to Yard
func MeterToYard(value float64) (float64, error) {
	return ConvertLength(value, "meter", "yard")
}

// Convert Meter to Foot
func MeterToFoot(value float64) (float64, error) {
	return ConvertLength(value, "meter", "foot")
}

// Convert Meter to Inch
func MeterToInch(value float64) (float64, error) {
	return ConvertLength(value, "meter", "inch")
}

// Convert Decimeter to Kilometer
func DecimeterToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "kilometer")
}

// Convert Decimeter to Hectometer
func DecimeterToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "hectometer")
}

// Convert Decimeter to Decameter
func DecimeterToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "decameter")
}

// Convert Decimeter to Meter
func DecimeterToMeter(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "meter")
}

// Convert Decimeter to Centimeter
func DecimeterToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "centimeter")
}

// Convert Decimeter to Millimeter
func DecimeterToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "millimeter")
}

// Convert Decimeter to Mile
func DecimeterToMile(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "mile")
}

// Convert Decimeter to Yard
func DecimeterToYard(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "yard")
}

// Convert Decimeter to Foot
func DecimeterToFoot(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "foot")
}

// Convert Decimeter to Inch
func DecimeterToInch(value float64) (float64, error) {
	return ConvertLength(value, "decimeter", "inch")
}

// Convert Centimeter to Kilometer
func CentimeterToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "kilometer")
}

// Convert Centimeter to Hectometer
func CentimeterToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "hectometer")
}

// Convert Centimeter to Decameter
func CentimeterToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "decameter")
}

// Convert Centimeter to Meter
func CentimeterToMeter(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "meter")
}

// Convert Centimeter to Decimeter
func CentimeterToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "decimeter")
}

// Convert Centimeter to Millimeter
func CentimeterToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "millimeter")
}

// Convert Centimeter to Mile
func CentimeterToMile(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "mile")
}

// Convert Centimeter to Yard
func CentimeterToYard(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "yard")
}

// Convert Centimeter to Foot
func CentimeterToFoot(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "foot")
}

// Convert Centimeter to Inch
func CentimeterToInch(value float64) (float64, error) {
	return ConvertLength(value, "centimeter", "inch")
}

// Convert Millimeter to Kilometer
func MillimeterToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "kilometer")
}

// Convert Millimeter to Hectometer
func MillimeterToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "hectometer")
}

// Convert Millimeter to Decameter
func MillimeterToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "decameter")
}

// Convert Millimeter to Meter
func MillimeterToMeter(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "meter")
}

// Convert Millimeter to Decimeter
func MillimeterToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "decimeter")
}

// Convert Millimeter to Centimeter
func MillimeterToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "centimeter")
}

// Convert Millimeter to Mile
func MillimeterToMile(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "mile")
}

// Convert Millimeter to Yard
func MillimeterToYard(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "yard")
}

// Convert Millimeter to Foot
func MillimeterToFoot(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "foot")
}

// Convert Millimeter to Inch
func MillimeterToInch(value float64) (float64, error) {
	return ConvertLength(value, "millimeter", "inch")
}

// Convert Mile to Kilometer
func MileToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "mile", "kilometer")
}

// Convert Mile to Hectometer
func MileToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "mile", "hectometer")
}

// Convert Mile to Decameter
func MileToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "mile", "decameter")
}

// Convert Mile to Meter
func MileToMeter(value float64) (float64, error) {
	return ConvertLength(value, "mile", "meter")
}

// Convert Mile to Decimeter
func MileToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "mile", "decimeter")
}

// Convert Mile to Centimeter
func MileToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "mile", "centimeter")
}

// Convert Mile to Millimeter
func MileToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "mile", "millimeter")
}

// Convert Mile to Yard
func MileToYard(value float64) (float64, error) {
	return ConvertLength(value, "mile", "yard")
}

// Convert Mile to Foot
func MileToFoot(value float64) (float64, error) {
	return ConvertLength(value, "mile", "foot")
}

// Convert Mile to Inch
func MileToInch(value float64) (float64, error) {
	return ConvertLength(value, "mile", "inch")
}

// Convert Yard to Kilometer
func YardToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "yard", "kilometer")
}

// Convert Yard to Hectometer
func YardToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "yard", "hectometer")
}

// Convert Yard to Decameter
func YardToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "yard", "decameter")
}

// Convert Yard to Meter
func YardToMeter(value float64) (float64, error) {
	return ConvertLength(value, "yard", "meter")
}

// Convert Yard to Decimeter
func YardToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "yard", "decimeter")
}

// Convert Yard to Centimeter
func YardToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "yard", "centimeter")
}

// Convert Yard to Millimeter
func YardToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "yard", "millimeter")
}

// Convert Yard to Mile
func YardToMile(value float64) (float64, error) {
	return ConvertLength(value, "yard", "mile")
}

// Convert Yard to Foot
func YardToFoot(value float64) (float64, error) {
	return ConvertLength(value, "yard", "foot")
}

// Convert Yard to Inch
func YardToInch(value float64) (float64, error) {
	return ConvertLength(value, "yard", "inch")
}

// Convert Foot to Kilometer
func FootToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "foot", "kilometer")
}

// Convert Foot to Hectometer
func FootToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "foot", "hectometer")
}

// Convert Foot to Decameter
func FootToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "foot", "decameter")
}

// Convert Foot to Meter
func FootToMeter(value float64) (float64, error) {
	return ConvertLength(value, "foot", "meter")
}

// Convert Foot to Decimeter
func FootToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "foot", "decimeter")
}

// Convert Foot to Centimeter
func FootToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "foot", "centimeter")
}

// Convert Foot to Millimeter
func FootToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "foot", "millimeter")
}

// Convert Foot to Mile
func FootToMile(value float64) (float64, error) {
	return ConvertLength(value, "foot", "mile")
}

// Convert Foot to Yard
func FootToYard(value float64) (float64, error) {
	return ConvertLength(value, "foot", "yard")
}

// Convert Foot to Inch
func FootToInch(value float64) (float64, error) {
	return ConvertLength(value, "foot", "inch")
}

// Convert Inch to Kilometer
func InchToKilometer(value float64) (float64, error) {
	return ConvertLength(value, "inch", "kilometer")
}

// Convert Inch to Hectometer
func InchToHectometer(value float64) (float64, error) {
	return ConvertLength(value, "inch", "hectometer")
}

// Convert Inch to Decameter
func InchToDecameter(value float64) (float64, error) {
	return ConvertLength(value, "inch", "decameter")
}

// Convert Inch to Meter
func InchToMeter(value float64) (float64, error) {
	return ConvertLength(value, "inch", "meter")
}

// Convert Inch to Decimeter
func InchToDecimeter(value float64) (float64, error) {
	return ConvertLength(value, "inch", "decimeter")
}

// Convert Inch to Centimeter
func InchToCentimeter(value float64) (float64, error) {
	return ConvertLength(value, "inch", "centimeter")
}

// Convert Inch to Millimeter
func InchToMillimeter(value float64) (float64, error) {
	return ConvertLength(value, "inch", "millimeter")
}

// Convert Inch to Mile
func InchToMile(value float64) (float64, error) {
	return ConvertLength(value, "inch", "mile")
}

// Convert Inch to Yard
func InchToYard(value float64) (float64, error) {
	return ConvertLength(value, "inch", "yard")
}

// Convert Inch to Foot
func InchToFoot(value float64) (float64, error) {
	return ConvertLength(value, "inch", "foot")
}
