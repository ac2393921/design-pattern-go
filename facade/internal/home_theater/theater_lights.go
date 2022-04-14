package hometheater

import (
	"fmt"
)

type TheaterLights struct {
	description string
}

func NewTheaterLights(description string) *TheaterLights {
	return &TheaterLights{
		description: description,
	}
}

func (t *TheaterLights) On() {
	fmt.Println(t.description + " on")
}

func (t *TheaterLights) Off() {
	fmt.Println(t.description + " off")
}

func (t *TheaterLights) Dim(level int) {
	fmt.Printf("%s dimming to %d %%", t.description, level)
}

func (t *TheaterLights) ToString() string {
	return t.description
}
