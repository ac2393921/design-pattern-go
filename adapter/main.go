package main

import (
	"fmt"

	d "github.com/ac2393921/design-pattern-go/adapter/internal/duck"
	t "github.com/ac2393921/design-pattern-go/adapter/internal/turkey"
)

func main() {
	duck := &d.MallardDuck{}
	turkey := &t.WildTurkey{}

	turkeyAdapter := d.NewTurkeyAdapter(turkey)

	fmt.Println("Turkeyの出力")
	turkey.Gobble()
	turkey.Fly()

	fmt.Println("Duckの出力")
	duck.Quack()
	duck.Fly()

	fmt.Println("TurkeyAdapterの出力")
	turkeyAdapter.Quack()
	turkeyAdapter.Fly()
}
