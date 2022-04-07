package pizza

import (
	"fmt"

	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/cheese"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/clams"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/dougth"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/pepperoni"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/sauce"
	"github.com/ac2393921/design-pattern-go/abstract_factory/internal/pizza_ingredient/veggies"
)

type PizzaInterface interface {
	Prepare()
	Bake()
	Cut()
	Box()
	SetName(string)
	GetName() string
}

type Pizza struct {
	name      string
	Dougth    dougth.Dough
	Sauce     sauce.Sauce
	Veggies   veggies.Veggies
	Cheese    cheese.Cheese
	Pepperoni pepperoni.Pepperoni
	Clam      clams.Clams
}

func (p *Pizza) Bake() {
	fmt.Println("350度で25分間焼く")
}

func (p *Pizza) Cut() {
	fmt.Println("ピザを扇形に切り分ける")
}

func (p *Pizza) Box() {
	fmt.Println("PizzaStoreの正式な箱にピザを入れる")
}

func (p *Pizza) SetName(name string) {
	p.name = name
}

func (p *Pizza) GetName() string {
	return p.name
}
