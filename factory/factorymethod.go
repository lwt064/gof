package factory

import "fmt"

// Mouse 鼠标
type Mouse interface {
	GetName() string
	GetPrice() float32
	GetDesc() string
}

// DellMouse 戴尔鼠标
type DellMouse struct {
	name  string
	price float32
}

func (a *DellMouse) GetName() string {
	return a.name
}

func (a *DellMouse) GetPrice() float32 {
	return a.price
}

func (a *DellMouse) GetDesc() string {
	return fmt.Sprintf("name[%s] price[%.2f]", a.name, a.price)
}

// HpMouse 惠普鼠标
type HpMouse struct {
	name  string
	price float32
}

func (b *HpMouse) GetName() string {
	return b.name
}

func (b *HpMouse) GetPrice() float32 {
	return b.price
}

func (b *HpMouse) GetDesc() string {
	return fmt.Sprintf("name[%s] price[%.2f]", b.name, b.price)
}

type MouseFactory interface {
	CreateMouse() Mouse
}

// FactoryDell 戴尔工厂
type DellFactory struct {
}

func (f *DellFactory) CreateMouse() Mouse {
	return &DellMouse{}
}

// FactoryHp 惠普工厂
type HpFactory struct {
}

func (f *HpFactory) CreateMouse() Mouse {
	return &HpMouse{}
}
