package main

import "fmt"

// 产品
type House struct {
	Door   string
	Window string
}

type HouseBuilder interface {
	BuildDoor(val string)
	BuildWindow(val string)
	GetHouse() *House
}

// 真正干活的
type Bao struct {
	house *House
}

func NewBao() *Bao {
	return &Bao{
		house: &House{},
	}
}

func (b *Bao) BuildDoor(val string) {
	b.house.Door = val
	fmt.Println("门建造成功")
}

func (b *Bao) BuildWindow(val string) {
	b.house.Window = val
	fmt.Println("窗建造成功")
}

func (b *Bao) GetHouse() *House {
	return b.house
}

// 管理者
type Boos struct {
	builder HouseBuilder
}

func NewBoos(bao *Bao) *Boos {
	return &Boos{
		builder: bao,
	}
}

// boos只需要得到建造的结果
func (b *Boos) GetHouse() *House {
	b.builder.BuildDoor("门")
	b.builder.BuildWindow("窗")
	return b.builder.GetHouse()
}

func main() {
	b := NewBao()

	boos := NewBoos(b)

	fmt.Println(boos.GetHouse())
}
