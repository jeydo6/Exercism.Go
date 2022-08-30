// Package weather forecast.
package weather

// CurrentCondition current condition.
var CurrentCondition string

// CurrentLocation current location.
var CurrentLocation string

// Forecast weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
