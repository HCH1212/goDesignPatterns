// 假设我们有一个饮料制作系统，支持制作咖啡和茶。咖啡和茶的制作过程有一些共同的步骤（如烧水、倒入杯子），也有一些不同的步骤（如冲泡咖啡、泡茶）。

package main

import "fmt"

// 模板
type Template interface {
	BoilWater()        // 烧水
	Brew()             // 冲泡
	AddSugar()         // 加糖
	HasAddSugar() bool // 是否加糖
}

type Coffee struct {
}

func (Coffee) BoilWater() {
	fmt.Println("烧水")
}

func (Coffee) Brew() {
	fmt.Println("冲泡")
}

func (Coffee) AddSugar() {
	fmt.Println("加糖")
}

func (Coffee) HasAddSugar() bool {
	return true
}

type Tea struct{}

func (Tea) BoilWater() {
	fmt.Println("烧水")
}

func (Tea) Brew() {
	fmt.Println("冲泡")
}

func (Tea) AddSugar() {
	fmt.Println("加糖")
}

func (Tea) HasAddSugar() bool {
	return false
}

// 模板方法
func MakeTemplate(tmp Template) {
	tmp.BoilWater()
	tmp.Brew()
	if tmp.HasAddSugar() {
		tmp.AddSugar()
	}
}

func main() {
	tea := Tea{}
	coffee := Coffee{}

	MakeTemplate(tea)

	MakeTemplate(coffee)
}
