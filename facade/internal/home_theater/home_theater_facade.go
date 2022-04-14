package hometheater

import "fmt"

type HomeTheaterFacade struct {
	amp       Amplifier
	tuner     Tuner
	player    StreamingPlayer
	cd        CdPlayer
	projector Projector
	lights    TheaterLights
	screen    Screen
	popper    PopcornPopper
}

func NewHomeTheaterFacade(
	amp Amplifier,
	tuner Tuner,
	player StreamingPlayer,
	cd CdPlayer,
	projector Projector,
	lights TheaterLights,
	screen Screen,
	popper PopcornPopper) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp:       amp,
		tuner:     tuner,
		player:    player,
		cd:        cd,
		projector: projector,
		lights:    lights,
		screen:    screen,
		popper:    popper,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.popper.On()
	h.popper.Pop()
	h.lights.Dim(10)
	h.screen.Down()
	h.projector.On()
	h.projector.WideScreenMode()
	h.amp.On()
	h.amp.SetStreamingPlayer(h.player)
	h.amp.SetVolume(5)
	h.player.On()
	h.player.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	h.popper.Off()
	h.lights.On()
	h.screen.Up()
	h.projector.Off()
	h.amp.Off()
	h.player.Stop()
	h.player.Off()
}

func (h *HomeTheaterFacade) ListenToRadio(frequency float32) {
	fmt.Println("Tuning in the airwaves...")
	h.tuner.On()
	h.tuner.SetFrequency(frequency)
	h.amp.On()
	h.amp.SetVolume(5)
	h.amp.SetTuner(h.tuner)
}

func (h *HomeTheaterFacade) EndRadio() {
	fmt.Println("Shutting down the tuner...")
	h.tuner.Off()
	h.amp.Off()
}
