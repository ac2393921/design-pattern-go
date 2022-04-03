package topping

import "github.com/ac2393921/design-pattern-go/decorator/internal/beverage"

type Soy struct {
	Beverage beverage.Beverage
}

func (s *Soy) GetDescription() string {
	return s.Beverage.GetDescription() + "、ソイ"
}

func (s *Soy) Cost() float32 {
	return s.Beverage.Cost() + .15
}
