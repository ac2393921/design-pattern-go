package turkey

import "fmt"

type WildTurkey struct{}

func (w *WildTurkey) Gobble() {
	fmt.Println("ゴロゴロ")
}

func (w *WildTurkey) Fly() {
	fmt.Println("短い距離を飛んでいます！")
}
