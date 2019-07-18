package factory

import "fmt"

type KeyBoard interface {
	GetName() string
	GetPrice() float32
	GetDesc() string
}

// DellKeyboard 戴尔键盘
type DellKeyboard struct {
	name  string
	price float32
}

func (a *DellKeyboard) GetName() string {
	return a.name
}

func (a *DellKeyboard) GetPrice() float32 {
	return a.price
}

func (a *DellKeyboard) GetDesc() string {
	return fmt.Sprintf("name[%s] price[%.2f]", a.name, a.price)
}

// HpKeyboard 惠普键盘
type HpKeyboard struct {
	name  string
	price float32
}

func (b *HpKeyboard) GetName() string {
	return b.name
}

func (b *HpKeyboard) GetPrice() float32 {
	return b.price
}

func (b *HpKeyboard) GetDesc() string {
	return fmt.Sprintf("name[%s] price[%.2f]", b.name, b.price)
}

// AbstractFactory, 抽象工厂
type AbstractFactory interface {
	CreateMouse() Mouse
	CreateKeyboard() KeyBoard
}

func (f *DellFactory) CreateKeyboard() KeyBoard {
	return &DellMouse{}
}

func (f *HpFactory) CreateKeyboard() KeyBoard {
	return &DellMouse{}
}
