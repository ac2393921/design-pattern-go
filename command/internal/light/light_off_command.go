package light

type LitghOffCommand struct {
	Light Light
}

func NewLigthOffCommand(light Light) *LitghOffCommand {
	return &LitghOffCommand{
		Light: light,
	}
}

func (l *LitghOffCommand) Execute() {
	l.Light.Off()
}
