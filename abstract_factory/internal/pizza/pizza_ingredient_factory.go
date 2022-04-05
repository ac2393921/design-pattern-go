package pizza

import (
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/cheese"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/clams"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/dougth"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/pepperoni"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/sauce"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/veggies"
)

type PizzaIngredientFactory interface {
	CreateDough() dougth.Dough
	CreateSouce() sauce.Sauce
	CreateCheese() cheese.Cheese
	CreateVeggies() []veggies.Veggies
	CreatePepperoni() pepperoni.Pepperoni
	CreateClam() clams.Clams
}
