package pizza

import (
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/cheese"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/clams"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/dougth"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/pepperoni"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/sauce"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/veggies"
)

type NYPizzaIngredientFactory struct{}

func (nypif *NYPizzaIngredientFactory) CreateDough() dougth.Dough {
	return &dougth.ThinCrustDough{}
}

func (nypif *NYPizzaIngredientFactory) CreateSauce() sauce.Sauce {
	return &sauce.MarinaraSauce{}
}

func (nypif *NYPizzaIngredientFactory) CreateCheese() cheese.Cheese {
	return &cheese.ReggianoCheese{}
}

func (nypif *NYPizzaIngredientFactory) CreateVeggies() []veggies.Veggies {
	veggies := []veggies.Veggies{&veggies.Garlic{}, &veggies.Onion{}, &veggies.Mushroom{}, &veggies.RedPepper{}}
	return veggies
}

func (nypif *NYPizzaIngredientFactory) CreatePepperoni() pepperoni.Pepperoni {
	return &pepperoni.SlicedPepperoni{}
}

func (nypif *NYPizzaIngredientFactory) CreateClam() clams.Clams {
	return &clams.FreshClams{}
}
