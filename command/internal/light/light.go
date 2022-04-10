package light

import "fmt"

type Light struct {
	Location string
}

func NewLight(location string) *Light {
	return &Light{
		Location: location,
	}
}

func (l *Light) On() {
	fmt.Printf("%s light is on", l.Location)
}

func (l *Light) Off() {
	fmt.Printf("%s light is off", l.Location)
}
