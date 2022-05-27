// Package weather something comments.
package weather

// CurrentCondition something comments.
var CurrentCondition string
// CurrentLocation something comments.
var CurrentLocation string

// Forecast something comments.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
