package observer

type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	wd := &WeatherData{}
	wd.observers = []Observer{}

	return wd
}

func (wd *WeatherData) RegisterObserver(observer Observer) {
	wd.observers = append(wd.observers, observer)
}

func (wd *WeatherData) RemoveObserver(observer Observer) {
	result := WeatherData{}
	for _, o := range wd.observers {
		if o != observer {
			result.observers = append(result.observers, o)
		}
	}
	wd.observers = result.observers
}

func (wd *WeatherData) NotifyObservers() {
	for _, o := range wd.observers {
		o.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}

func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.MeasurementsChanged()
}
