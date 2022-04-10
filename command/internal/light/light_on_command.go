package light

type LigthONCommand struct {
	Light Light
}

func NewLigthOnCommand(light Light) *LigthONCommand {
	return &LigthONCommand{
		Light: light,
	}
}

func (l *LigthONCommand) Execute() {
	l.Light.On()
}
