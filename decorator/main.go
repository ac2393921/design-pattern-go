package main

import (
	"fmt"

	"github.com/ac2393921/design-pattern-go/decorator/internal/beverage"
	"github.com/ac2393921/design-pattern-go/decorator/internal/topping"
)

func main() {
	espresso := &beverage.Espresso{}
	espresso_beverage := beverage.Beverage(espresso)
	fmt.Printf("%s $%.2f\n", espresso_beverage.GetDescription(), espresso_beverage.Cost())

	darkRoast := &beverage.DarkRoast{}
	darkRoast_beverage := beverage.Beverage(darkRoast)

	fmt.Printf("%s $%.2f\n", darkRoast_beverage.GetDescription(), darkRoast_beverage.Cost())

	singleMocha := &topping.Mocha{
		Beverage: darkRoast_beverage,
	}
	fmt.Printf("%s $%.2f\n", singleMocha.GetDescription(), singleMocha.Cost())

	doubleMocha := &topping.Mocha{
		Beverage: singleMocha,
	}
	fmt.Printf("%s $%.2f\n", doubleMocha.GetDescription(), doubleMocha.Cost())

	doubleMochaWhip := &topping.Whip{
		Beverage: doubleMocha,
	}
	fmt.Printf("%s $%.2f\n", doubleMochaWhip.GetDescription(), doubleMochaWhip.Cost())

	soyMochaWhipHouseBlend := &topping.Whip{
		Beverage: &topping.Mocha{
			Beverage: &topping.Soy{
				Beverage: &beverage.HouseBlend{},
			},
		},
	}
	fmt.Printf("%s $%.2f\n", soyMochaWhipHouseBlend.GetDescription(), soyMochaWhipHouseBlend.Cost())
}
