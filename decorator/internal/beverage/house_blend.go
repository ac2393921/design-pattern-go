package beverage

type HouseBlend struct{}

func (h *HouseBlend) GetDescription() string {
	return "ハウスブレンドコーヒー"
}

func (h *HouseBlend) Cost() float32 {
	return .89
}
