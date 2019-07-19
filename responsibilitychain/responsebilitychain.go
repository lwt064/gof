package responsibilitychain

type Handler func(input map[string]interface{}, c *Context)

type Context struct {
	Handlers []Handler
	Data     map[string]interface{} // 中间数据
	Cur      int                    // 当前handler下标
	Finished bool                   // 是否完成处理
}

type engine struct {
	c *Context
}

func (e *engine) AddHandler(h ...Handler) {
	c := e.c
	c.Handlers = append(c.Handlers, h...)
}

func (e *engine) HandleRequest(input map[string]interface{}) {
	c := e.c
	for i, handler := range c.Handlers {
		c.Cur = i
		if c.Finished {
			break
		}
		handler(input, c)
	}
}
