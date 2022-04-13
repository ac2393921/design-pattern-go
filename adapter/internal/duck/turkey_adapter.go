package duck

import "github.com/ac2393921/design-pattern-go/adapter/internal/turkey"

type TurkeyAdapter struct {
	Turkey turkey.Turkey
}

func NewTurkeyAdapter(turkey turkey.Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{
		Turkey: turkey,
	}
}

func (t *TurkeyAdapter) Quack() {
	t.Turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.Turkey.Fly()
	}
}
