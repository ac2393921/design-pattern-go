package remote

import "fmt"

type NoCommand struct{}

func NewNoCommand() *NoCommand {
	return &NoCommand{}
}

func (n *NoCommand) Execute() {
	fmt.Print("no command set")
}
