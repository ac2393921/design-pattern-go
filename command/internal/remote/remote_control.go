package remote

import (
	"fmt"
	"reflect"
)

type RemoteControl struct {
	OnCommand  []Command
	OffCommand []Command
}

func NewRemoteControl() *RemoteControl {
	onCommand := make([]Command, 7)
	offCommand := make([]Command, 7)

	for i := 0; i < 7; i++ {
		onCommand[i] = NewNoCommand()
		offCommand[i] = NewNoCommand()
	}

	return &RemoteControl{
		OnCommand:  onCommand,
		OffCommand: offCommand,
	}
}

func (r *RemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	r.OnCommand[slot] = onCommand
	r.OffCommand[slot] = offCommand
}

func (r *RemoteControl) OnButtonWasPushed(slot int) {
	fmt.Print("\n")
	r.OnCommand[slot].Execute()
}

func (r *RemoteControl) OffButtonWasPushed(slot int) {
	fmt.Print("\n")
	r.OffCommand[slot].Execute()
}

func (r *RemoteControl) ToString() string {
	s := "\n------ Remote Control ------\n"

	for i := 0; i < 7; i++ {
		if r.OnCommand[i] == nil {
			continue
		}

		onClass := r.getClassName(r.OnCommand[i])
		offClass := r.getClassName(r.OffCommand[i])
		s += fmt.Sprintf("[slot %d] %s   %s\n", i, onClass, offClass)
	}
	s += "-----------------------------\n"

	return s
}

func (r *RemoteControl) getClassName(command Command) string {
	if t := reflect.TypeOf(command); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
