package main

import (
	"github.com/ac2393921/design-pattern-go/observer/internal/observer"
)

func main() {
	wd := observer.NewWeatherData()

	ccd := observer.NewCurrentConditionsDisplay(wd)
	o := observer.Observer(ccd)
	// fmt.Printf("%T\n", ccd)
	// fmt.Printf("%T\n", o)
	wd.RegisterObserver(o)

	wd.SetMeasurements(80, 65, 30.4)
	wd.SetMeasurements(82, 70, 29.2)
	wd.SetMeasurements(77, 90, 29.2)
}
