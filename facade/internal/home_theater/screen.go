package hometheater

import (
	"fmt"
)

type Screen struct {
	description string
}

func NewScreen(description string) *Screen {
	return &Screen{
		description: description,
	}
}

func (s *Screen) Up() {
	fmt.Println(s.description + " going up")
}

func (s *Screen) Down() {
	fmt.Println(s.description + " going down")
}

func (s *Screen) ToString() string {
	return s.description
}
