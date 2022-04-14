package hometheater

import "fmt"

type CdPlayer struct {
	description  string
	Amplifier    Amplifier
	title        string
	currentTrack int
}

func NewCdPlayer(description string, amplifier Amplifier) *CdPlayer {
	return &CdPlayer{
		description:  description,
		Amplifier:    amplifier,
		title:        "",
		currentTrack: 0,
	}
}

func (c *CdPlayer) On() {
	fmt.Println(c.description + " on")
}

func (c *CdPlayer) Off() {
	fmt.Println(c.description + " off")
}

func (c *CdPlayer) Eject() {
	c.title = ""
	fmt.Println(c.description + " eject")
}

func (c *CdPlayer) PlayByTitle(title string) {
	c.title = title
	c.currentTrack = 0
	fmt.Println(c.description + " playing \"" + title + "\"")
}

func (c *CdPlayer) PlayByTrack(track int) {
	if c.title == "" {
		fmt.Printf("%s can't play track %d, no cd inserted", c.description, c.currentTrack)
	} else {
		c.currentTrack = track
		fmt.Printf("%s playing track %d", c.description, c.currentTrack)
	}
}

func (c *CdPlayer) Stop() {
	c.currentTrack = 0
	fmt.Println(c.description + " stop")
}

func (c *CdPlayer) Pause() {
	fmt.Println(c.description + " pause")
}

func (c *CdPlayer) ToString() string {
	return c.description
}
