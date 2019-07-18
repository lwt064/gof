package factory

import (
	"fmt"
)

// Item 产品
type Item interface {
	GetName() string
	GetPrice() float32
	GetDesc() string
}

// ItemA 产品A
type ItemA struct {
	name  string
	price float32
}

func (a *ItemA) GetName() string {
	return a.name
}

func (a *ItemA) GetPrice() float32 {
	return a.price
}

func (a *ItemA) GetDesc() string {
	return fmt.Sprintf("name[%s] price[%.2f]", a.name, a.price)
}

// ItemB 产品B
type ItemB struct {
	name  string
	price float32
}

func (b *ItemB) GetName() string {
	return b.name
}

func (b *ItemB) GetPrice() float32 {
	return b.price
}

func (b *ItemB) GetDesc() string {
	return fmt.Sprintf("name[%s] price[%.2f]", b.name, b.price)
}

type SimpleFactory struct {
}

func (f *SimpleFactory) CreateItem(name string) Item {
	if name == "A" {
		return &ItemA{}
	} else if name == "B" {
		return &ItemB{}
	} else {
		return nil
	}
}
