// Package weather provides the curren location's current weather condition.
package weather

// CurrentCondition stores a string variable which contains the current condition of an area.
var CurrentCondition string

// CurrentLocation stores a string variable which contains the current location.
var CurrentLocation string

// Forecast provides the current location and current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
