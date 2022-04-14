package hometheater

import "fmt"

type Tuner struct {
	description string
}

func NewTuner(description string) *Tuner {
	return &Tuner{
		description: description,
	}
}

func (t *Tuner) On() {
	fmt.Println(t.description + " on")
}

func (t *Tuner) Off() {
	fmt.Println(t.description + " off")
}

func (t *Tuner) SetFrequency(frequency float32) {
	fmt.Printf("%s setting frequency to %.2f", t.description, frequency)
}

func (t *Tuner) SetAm() {
	fmt.Println(t.description + " setting AM mode")
}

func (t *Tuner) SetFm() {
	fmt.Println(t.description + " setting FM mode")
}

func (t *Tuner) ToString() string {
	return t.description
}
