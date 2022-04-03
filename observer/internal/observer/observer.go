package observer

type Observer interface {
	Update(temperature, humidity, pressure float64)
}
