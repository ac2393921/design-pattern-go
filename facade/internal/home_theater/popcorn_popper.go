package hometheater

import "fmt"

type PopcornPopper struct {
	description string
}

func NewPopcornPopper(description string) *PopcornPopper {
	return &PopcornPopper{
		description: description,
	}
}

func (p *PopcornPopper) On() {
	fmt.Println(p.description + " on")
}

func (p *PopcornPopper) Off() {
	fmt.Println(p.description + " off")
}

func (p *PopcornPopper) Pop() {
	fmt.Println(p.description + " popping popcorn!")
}

func (p *PopcornPopper) ToString() string {
	return p.description
}
