package hometheater

import "fmt"

type Projector struct {
	description string
	player      StreamingPlayer
}

func NewProjector(description string, player StreamingPlayer) *Projector {
	return &Projector{
		description: description,
		player:      player,
	}
}

func (p *Projector) On() {
	fmt.Println(p.description + " on")
}

func (p *Projector) Off() {
	fmt.Println(p.description + " off")
}

func (p *Projector) WideScreenMode() {
	fmt.Println(p.description + " in widescreen mode (16x9 aspect ratio)")
}

func (p *Projector) TvMode() {
	fmt.Println(p.description + " in tv mode (4x3 aspect ratio)")
}

func (p *Projector) ToString() string {
	return p.description
}
