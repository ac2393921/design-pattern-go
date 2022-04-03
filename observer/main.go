package main

import (
	"github.com/ac2393921/design-pattern-go/internal/observer"
)

func main() {
	wd := observer.WeatherData.NewWeatherData()

	currentConditionDisplay := NewCurrentConditionDisplay()
	wd.RegisterObserver(currentConditionDisplay)

	wd.setMeasurements(80, 65, 30.4)
	wd.setMeasurements(82, 70, 29.2)
	wd.setMeasurements(77, 90, 29.2)
}
