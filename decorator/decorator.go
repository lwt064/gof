package decorator

import (
	"fmt"
	"time"
)

type UserLogic func(input map[string]interface{}) string

// 装饰模式，不改变对象结构，但动态扩展了它的功能。
func Decorate(f UserLogic) UserLogic {
	return func(input map[string]interface{}) string {
		fmt.Println("start on ", time.Now())
		output := f(input)
		fmt.Println("end on ", time.Now())
		return output
	}
}
