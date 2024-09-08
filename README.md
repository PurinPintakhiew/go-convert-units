# go-convert-units ✨
go-convert-units is a Go package designed to simplify conversions between units of measurement, including weight, length, and temperature. It provides a user-friendly interface for developers to convert values ​​between different units.

## Install 🛠️
```
go get github.com/PurinPintakhiew/go-convert-units
```
## How to use💡
```
package main

import (
	"fmt"

	"github.com/PurinPintakhiew/go-convert-units"
)

func main() {
	// --- Convert length units
	calcLength := convertunits.ConvertLength(1, "kilometer", "meter")
	fmt.Printf("length: %f\n", calcLength) // length: 1000

	// --- Convert weight units
	calcWeight := convertunits.ConvertWeight(1, "kilogram", "gram")
	fmt.Printf("weight: %f\n", calcWeight) // weight: 1000

	// --- Convert temperature units
	calcTemp := convertunits.ConvertTemperature(0, "celsius", "fahrenheit")
	fmt.Printf("temp: %f\n", calcTemp) // temp: 32
}
```
