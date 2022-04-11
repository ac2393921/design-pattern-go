package ceilingfun

type CeilingFunOnCommand struct {
	CeilingFun CeilingFun
}

func NewCeilingFunOnCommand(ceilingFun CeilingFun) *CeilingFunOnCommand {
	return &CeilingFunOnCommand{
		CeilingFun: ceilingFun,
	}
}

func (c *CeilingFunOnCommand) Execute() {
	c.CeilingFun.High()
}
