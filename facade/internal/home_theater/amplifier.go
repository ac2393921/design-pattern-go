package hometheater

import "fmt"

type Amplifier struct {
	description string
	tuner       Tuner
}

func NewAmplifier(description string) *Amplifier {
	return &Amplifier{
		description: description,
	}
}

func (a *Amplifier) On() {
	fmt.Println(a.description + " on")
}

func (a *Amplifier) Off() {
	fmt.Println(a.description + " off")
}

func (a *Amplifier) SetStereoSound() {
	fmt.Println(a.description + " stereo mode on")
}

func (a *Amplifier) SetSurroundSound() {
	fmt.Println(a.description + " surround sound on (5 speakers, 1 subwoofer)")
}

func (a *Amplifier) SetVolume(level int) {
	fmt.Printf("%s setting volume to %d", a.description, level)
}

func (a *Amplifier) SetTuner(tuner Tuner) {
	fmt.Printf("%s setting volume to %s", a.description, tuner)
	a.tuner = tuner
}

func (a *Amplifier) SetStreamingPlayer(player StreamingPlayer) {
	fmt.Println(a.description + " setting Streaming player to" + player.ToString())
}

func (a *Amplifier) ToString() string {
	return a.description
}
