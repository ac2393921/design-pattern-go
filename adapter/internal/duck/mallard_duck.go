package duck

import "fmt"

type MallardDuck struct{}

func (m *MallardDuck) Quack() {
	fmt.Println("ガーガー")
}

func (m *MallardDuck) Fly() {
	fmt.Println("飛んでいます！")
}
