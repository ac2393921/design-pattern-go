package beverage

type Espresso struct{}

func (e *Espresso) GetDescription() string {
	return "エスプレッソ"
}

func (e *Espresso) Cost() float32 {
	return 1.99
}
