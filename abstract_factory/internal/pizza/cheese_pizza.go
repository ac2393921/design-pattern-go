package pizza

import (
	"fmt"
)

type CheesePizza struct {
	Pizza
	factory PizzaIngredientFactory
}

func NewCheesePizza(factory PizzaIngredientFactory) *CheesePizza {
	cp := &CheesePizza{
		Pizza:   Pizza{},
		factory: factory,
	}
	return cp
}

func (cp *CheesePizza) Prepare() {
	fmt.Println(cp.name + "γδΈε¦η")
	cp.Dougth = cp.factory.CreateDough()
	cp.Sauce = cp.factory.CreateSauce()
	cp.Cheese = cp.factory.CreateCheese()
	cp.Clam = cp.factory.CreateClam()
}
