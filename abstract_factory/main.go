package main

import (
	"fmt"

	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_store"
)

func main() {
	nyPizzaStore := pizza_store.NewNYPizzaStore()
	p := nyPizzaStore.OrderPizza("cheese")
	fmt.Printf("You take out a %s", p.GetName())
}
