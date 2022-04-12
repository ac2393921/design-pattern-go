package stereo

type StereoOffCommand struct {
	Stereo Stereo
}

func NewStereoOffCommand(stereo Stereo) *StereoOffCommand {
	return &StereoOffCommand{
		Stereo: stereo,
	}
}

func (s *StereoOffCommand) Execute() {
	s.Stereo.Off()
}
