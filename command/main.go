package main

import (
	"fmt"

	ceilingfun "github.com/ac2393921/design-pattern-go/command/internal/ceiling_fun"
	"github.com/ac2393921/design-pattern-go/command/internal/light"
	"github.com/ac2393921/design-pattern-go/command/internal/remote"
	st "github.com/ac2393921/design-pattern-go/command/internal/stereo"
)

func main() {
	remoteControl := remote.NewRemoteControl()

	livingRoomLight := light.NewLight("Living Room")
	kitchenLight := light.NewLight("Kitchen Room")
	ceilingFun := ceilingfun.NewCeilingFun("Living Room")
	stereo := st.NewStereo("Living Room")

	livingRoomLightOn := light.NewLigthOnCommand(*livingRoomLight)
	livingRoomLightOff := light.NewLigthOffCommand(*livingRoomLight)

	kitchenLightOn := light.NewLigthOnCommand(*kitchenLight)
	kitchenLightOff := light.NewLigthOffCommand(*kitchenLight)

	ceilingfunOn := ceilingfun.NewCeilingFunOnCommand(*ceilingFun)
	ceilingfunOff := ceilingfun.NewCeilingFunOffCommand(*ceilingFun)

	stereoOn := st.NewStereoOnCommand(*stereo)
	stereoOff := st.NewStereoOffCommand(*stereo)

	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.SetCommand(2, ceilingfunOn, ceilingfunOff)
	remoteControl.SetCommand(3, stereoOn, stereoOff)

	fmt.Print(remoteControl)

	for i := 0; i < 7; i++ {
		remoteControl.OnButtonWasPushed(i)
		remoteControl.OffButtonWasPushed(i)
	}
}
