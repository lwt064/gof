package proxy

import "fmt"

// 用户找工作

type Company interface {
	Employ(uid int)
}

type CompanyX struct {
}

func (x *CompanyX) Employ(uid int) {
	fmt.Println("cons! you was employed.")
}

type Proxy struct {
	c *CompanyX
}

// 可以控制对真实对象的访问（比如条件检测）或实现一些附加操作，对扩展开放，对修改封闭
func (p *Proxy) Employ(uid int) {
	p.c.Employ(uid)
	fmt.Println("uid[%d] was employed.", uid)
}
