package topping

import "github.com/ac2393921/design-pattern-go/decorator/internal/beverage"

type Milk struct {
	Beverage beverage.Beverage
}

func (m *Milk) GetDescription() string {
	return m.Beverage.GetDescription() + "、ミルク"
}

func (m *Milk) Cost() float32 {
	return m.Beverage.Cost() + .1
}
