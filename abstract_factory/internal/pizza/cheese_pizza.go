package pizza

import (
	"fmt"
)

type CheesePizza struct {
	Pizza
	factory PizzaIngredientFactory
}

func NewCheesePizza(pif PizzaIngredientFactory) *CheesePizza {
	cp := &CheesePizza{
		Pizza:   Pizza{},
		factory: pif,
	}
	return cp
}

func (cp *CheesePizza) Prepare() {
	fmt.Println(cp.name + "を下処理")
	cp.Dougth = cp.factory.CreateDough()
	cp.Sauce = cp.factory.CreateSauce()
	cp.Cheese = cp.factory.CreateCheese()
	cp.Clam = cp.factory.CreateClam()
}
