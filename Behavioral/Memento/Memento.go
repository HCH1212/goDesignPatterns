package main

import "fmt"

type Node interface {
	SetState(state string)
	GetState() string
}

type TextNode struct {
	state string // 状态
}

func (t *TextNode) SetState(state string) {
	t.state = state
}

func (t *TextNode) GetState() string {
	return t.state
}

func (t *TextNode) Save() Memento {
	return &TextMemento{
		state: t.state,
	}
}

// 备忘录
type Memento interface {
	GetState() string
}

// 文本备忘录
type TextMemento struct {
	state string
}

func (t *TextMemento) GetState() string {
	return t.state
}

// 管理者
type Manage struct {
	states []Memento
}

func (m *Manage) Save(t Memento) {
	m.states = append(m.states, t)
}

func (m *Manage) Back(index int) Memento {
	return m.states[index]
}

func main() {
	manage := Manage{}
	text := TextNode{}

	text.SetState("1")
	manage.Save(text.Save())

	text.SetState("2")
	manage.Save(text.Save())

	fmt.Println(text.GetState())
	fmt.Println(manage.Back(0).GetState())
}
