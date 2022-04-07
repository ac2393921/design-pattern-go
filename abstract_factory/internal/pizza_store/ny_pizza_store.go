package pizza_store

import (
	"fmt"

	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza"
)

type NYPizzaStore struct {
	PizzaStore
	ingredientFactory pizza.PizzaIngredientFactory
}

func NewNYPizzaStore() *NYPizzaStore {
	nyps := &NYPizzaStore{
		PizzaStore:        PizzaStore{},
		ingredientFactory: &pizza.NYPizzaIngredientFactory{},
	}
	nyps.PizzaStore.createPizza = nyps.createPizza

	return nyps
}

func (nyps *NYPizzaStore) createPizza(item string) (pizza.PizzaInterface, error) {
	switch item {
	case "cheese":
		p := pizza.NewCheesePizza(nyps.ingredientFactory)
		p.SetName("New York Style Cheese Pizza")
		return p, nil
	}

	return nil, fmt.Errorf("invalid piiza type")
}
