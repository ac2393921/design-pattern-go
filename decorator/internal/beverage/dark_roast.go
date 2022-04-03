package beverage

type DarkRoast struct{}

func (d *DarkRoast) GetDescription() string {
	return "ダークローストコーヒー"
}

func (d *DarkRoast) Cost() float32 {
	return 0.99
}
