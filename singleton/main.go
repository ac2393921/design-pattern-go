package main

import (
	"github.com/ac2393921/design-pattern-go/singleton/internal/chocolate_boiler"
)

func main() {
	ch := make(chan interface{})
	go run(ch)
	go run(ch)
	go run(ch)
	<-ch
}

func run(ch chan interface{}) {
	ch <- chocolate_boiler.GetInstance()
}
