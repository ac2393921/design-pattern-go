package beverage

type Decaf struct{}

func (d *Decaf) GetDescription() string {
	return "デカフェ"
}

func (d *Decaf) Cost() float32 {
	return 1.05
}
