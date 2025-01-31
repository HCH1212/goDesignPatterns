// 比如开关的状态

package main

import "fmt"

// 状态接口
type State interface {
	Switch(context *Context)
}

// 状态上下文
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Switch() {
	c.state.Switch(c)
}

// 两个状态
type OnState struct{}

func (OnState) Switch(context *Context) {
	fmt.Println("开关关闭")
	context.SetState(&OffState{})
}

type OffState struct{}

func (OffState) Switch(context *Context) {
	fmt.Println("开关打开")
	context.SetState(&OnState{})
}

func main() {
	c := &Context{
		state: &OffState{}, // 默认开关关闭
	}
	c.Switch()
	c.Switch()
	c.Switch()
}
