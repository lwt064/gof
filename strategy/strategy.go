package strategy

import "fmt"

type Strategy interface {
	Do()
}

type Caller struct {
	stg Strategy
}

func (c *Caller) SetStrategy(s Strategy) {
	c.stg = s
}

func (c *Caller) DoSth() {
	fmt.Println("use stg x begin")
	s := c.stg
	s.Do()
	fmt.Println("use stg x end")
}

type StrategyA struct{}

func (a *StrategyA) Do() {
	fmt.Println("A")
}

type StrategB struct{}

func (a *StrategB) Do() {
	fmt.Println("B")
}

type StrategyC struct{}

func (a *StrategyC) Do() {
	fmt.Println("C")
}
