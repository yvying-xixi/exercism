// Package weather provides information about the current weather conditions
// and location.
package weather

var (
	// CurrentCondition stores the current weather condition.
	CurrentCondition string

	// CurrentLocation stores the location for the current weather condition.
	CurrentLocation string
)

// Forecast returns the current weather condition for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
