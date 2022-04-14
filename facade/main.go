package main

import hometheater "github.com/ac2393921/design-pattern-go/facade/internal/home_theater"

func main() {
	amp := hometheater.NewAmplifier("Amplifier")
	tuner := hometheater.NewTuner("AM/FM Tuner")
	player := hometheater.NewStreamingPlayer("Streaming Player", *amp)
	cd := hometheater.NewCdPlayer("CD Player", *amp)
	projector := hometheater.NewProjector("Projector", *player)
	lights := hometheater.NewTheaterLights("Theater Ceiling Lights")
	screen := hometheater.NewScreen("Theater Screen")
	popper := hometheater.NewPopcornPopper("Popcorn Popper")

	homeTheater := hometheater.NewHomeTheaterFacade(*amp, *tuner, *player, *cd, *projector, *lights, *screen, *popper)

	homeTheater.WatchMovie("Raiders of the Lost Ark")
	homeTheater.EndMovie()
}
