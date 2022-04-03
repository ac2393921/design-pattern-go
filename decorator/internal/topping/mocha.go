package topping

import "github.com/ac2393921/design-pattern-go/decorator/internal/beverage"

type Mocha struct {
	Beverage beverage.Beverage
}

func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + "、モカ"
}

func (m *Mocha) Cost() float32 {
	return m.Beverage.Cost() + .20
}
