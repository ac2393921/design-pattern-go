package pizza_store

import (
	"fmt"
	"runtime/debug"

	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza"
)

type PizzaStore struct {
	createPizza func(pizzaType string) (pizza.PizzaInterface, error)
}

type PizzaStoreInterface interface {
	OrderPizza(string) pizza.PizzaInterface
}

func (ps *PizzaStore) OrderPizza(item string) pizza.PizzaInterface {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	p, err := ps.createPizza(item)

	if err != nil {
		return nil
	}

	fmt.Printf("--- Making a %s ---", p.GetName())
	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()

	return p
}
