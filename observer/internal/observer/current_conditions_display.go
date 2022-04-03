package observer

import "fmt"

type CurrentConditionsDisplay struct {
	subject     *WeatherData
	temperature float64
	humidity    float64
	pressure    float64
}

func NewCurrentConditionsDisplay(wd *WeatherData) *CurrentConditionsDisplay {
	ccd := &CurrentConditionsDisplay{}
	ccd.subject = wd

	return ccd
}

func (ccd *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.pressure = pressure

	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("現在の気象状況:\n温度%g度\n湿度%g%%", ccd.temperature, ccd.humidity)
}
