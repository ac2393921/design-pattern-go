package ceilingfun

import "fmt"

var HIGH = 2
var MEDIUM = 2
var LOW = 0

type CeilingFun struct {
	Location string
	level    int
}

func NewCeilingFun(location string) *CeilingFun {
	return &CeilingFun{
		Location: location,
	}
}

func (c *CeilingFun) High() {
	c.level = HIGH
	fmt.Print(c.Location + " ceiling fan is on high")
}

func (c *CeilingFun) Medium() {
	c.level = MEDIUM
	fmt.Print(c.Location + " ceiling fan is on medium")
}

func (c *CeilingFun) Low() {
	c.level = LOW
	fmt.Print(c.Location + " ceiling fan is on low")
}

func (c *CeilingFun) Off() {
	c.level = 0
	fmt.Print(c.Location + " ceiling fan is off")
}

func (c *CeilingFun) GetSpeed() int {
	return c.level
}
