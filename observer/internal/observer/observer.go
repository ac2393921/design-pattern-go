package observer

type Observer interface {
	update(temperature, humidity, pressure float64)
}
