package adapter

import "fmt"

type Duck interface {
	quack() // 呱呱
	fly()   // 飞行
}

type Turkey interface {
	gobble() // 咯咯
	fly()    // 飞
}

type MallardDuck struct{}

func (m *MallardDuck) quack() {
	fmt.Println("Quack...")
}

func (m *MallardDuck) fly() {
	fmt.Println("flying...")
}

type WildTurkey struct{}

func (w *WildTurkey) gobble() {
	fmt.Println("Gobble...")
}

func (w *WildTurkey) fly() {
	fmt.Println("short flying...")
}

type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey}
}

func (ta *TurkeyAdapter) quack() {
	ta.turkey.gobble()
}

func (ta *TurkeyAdapter) fly() {
	for i := 0; i < 5; i++ {
		ta.turkey.fly()
	}
}
