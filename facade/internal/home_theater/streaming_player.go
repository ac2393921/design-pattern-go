package hometheater

import "fmt"

type StreamingPlayer struct {
	description    string
	amplifier      Amplifier
	movie          string
	currentChapter int
}

func NewStreamingPlayer(description string, amplifier Amplifier) *StreamingPlayer {
	return &StreamingPlayer{
		description:    description,
		amplifier:      amplifier,
		movie:          "",
		currentChapter: 0,
	}
}

func (s *StreamingPlayer) On() {
	fmt.Println(s.description + " on")
}

func (s *StreamingPlayer) Off() {
	fmt.Println(s.description + " off")
}

func (s *StreamingPlayer) Play(movie string) {
	s.movie = movie
	s.currentChapter = 0
	fmt.Printf("%s playing \"%s\"", s.description, s.movie)
}

func (s *StreamingPlayer) PlayByChapter(chapter int) {
	if s.movie == "" {
		fmt.Printf("%s can't play chapter %d no movie selected", s.movie, chapter)
	} else {
		s.currentChapter = chapter
		fmt.Printf("%s playing chapter \"%d\"", s.description, chapter)
	}
}

func (s *StreamingPlayer) Stop() {
	s.currentChapter = 0
	fmt.Println(s.description + " stop")
}

func (s *StreamingPlayer) Pause() {
	fmt.Println(s.description + " pause")
}

func (s *StreamingPlayer) SetTwoChannelAudio() {
	fmt.Println(s.description + " set two channel audio")
}

func (s *StreamingPlayer) SetSurroundAudio() {
	fmt.Println(s.description + " set surround audio")
}

func (s *StreamingPlayer) ToString() string {
	return s.description
}
