package ceilingfun

type CeilingFunOffCommand struct {
	CeilingFun CeilingFun
}

func NewCeilingFunOffCommand(ceilingFun CeilingFun) *CeilingFunOffCommand {
	return &CeilingFunOffCommand{
		CeilingFun: ceilingFun,
	}
}

func (c *CeilingFunOffCommand) Execute() {
	c.CeilingFun.Off()
}
