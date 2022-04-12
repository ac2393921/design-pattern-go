package stereo

import "fmt"

type Stereo struct {
	Location string
}

func NewStereo(location string) *Stereo {
	return &Stereo{
		Location: location,
	}
}

func (s *Stereo) On() {
	fmt.Print(s.Location + " stereo is on")
}

func (s *Stereo) Off() {
	fmt.Print(s.Location + " stereo is off")
}

func (s *Stereo) SetCD() {
	fmt.Print(s.Location + " set for CD input")
}

func (s *Stereo) SetDVD() {
	fmt.Print(s.Location + " set for DVD input")
}

func (s *Stereo) SetRadio() {
	fmt.Print(s.Location + " set for Radio input")
}

func (s *Stereo) SetVolume(volume int) {
	fmt.Print(s.Location+" stereo volume set to", volume)
}
