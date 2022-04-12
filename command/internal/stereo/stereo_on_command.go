package stereo

type StereoOnCommand struct {
	Stereo Stereo
}

func NewStereoOnCommand(stereo Stereo) *StereoOnCommand {
	return &StereoOnCommand{
		Stereo: stereo,
	}
}

func (s *StereoOnCommand) Execute() {
	s.Stereo.On()
}
