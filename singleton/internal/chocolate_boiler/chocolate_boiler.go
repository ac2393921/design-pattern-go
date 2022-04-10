package chocolate_boiler

import "fmt"

type ChocolateBoiler struct {
	Empty  bool
	Boiled bool
}

var sharedInstance *ChocolateBoiler = newChocolateBoiler()

func newChocolateBoiler() *ChocolateBoiler {
	fmt.Println("create new instance.")
	return &ChocolateBoiler{
		Empty:  false,
		Boiled: false,
	}
}

func GetInstance() *ChocolateBoiler {
	fmt.Println("get instance.")
	return sharedInstance
}

func (cb *ChocolateBoiler) Fill() {
	if cb.isEmpty() {
		cb.Empty = false
		cb.Boiled = false
	}
}

func (cb *ChocolateBoiler) Drain() {
	if !cb.isEmpty() && cb.isBoiled() {
		cb.Empty = true
	}
}

func (cb *ChocolateBoiler) Boil() {
	if !cb.isEmpty() && !cb.isBoiled() {
		cb.Boiled = true
	}
}

func (cb *ChocolateBoiler) isEmpty() bool {
	return cb.Empty
}

func (cb *ChocolateBoiler) isBoiled() bool {
	return cb.Boiled
}
