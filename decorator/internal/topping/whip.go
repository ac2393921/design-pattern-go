package topping

import "github.com/ac2393921/design-pattern-go/decorator/internal/beverage"

type Whip struct {
	Beverage beverage.Beverage
}

func (w *Whip) GetDescription() string {
	return w.Beverage.GetDescription() + "、ウィップ"
}

func (w *Whip) Cost() float32 {
	return w.Beverage.Cost() + .1
}
