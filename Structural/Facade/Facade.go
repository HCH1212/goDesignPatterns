package main

import "fmt"

// 库存
type Inventory struct{}

func (Inventory) Deduction() {
	fmt.Println("扣库存")
}

// 支付
type Pay struct{}

func (Pay) Pay() {
	fmt.Println("支付")
}

// 物流
type Logistics struct{}

func (Logistics) SendOutGoods() {
	fmt.Println("发货")
}

// 自己定义一个下单, 简化
type Order struct {
	i *Inventory
	p *Pay
	l *Logistics
}

func NewOrder() *Order {
	return &Order{
		i: &Inventory{},
		p: &Pay{},
		l: &Logistics{},
	}
}

func (o Order) Place() {
	o.i.Deduction()
	o.p.Pay()
	o.l.SendOutGoods()
}

func main() {
	//i := Inventory{}
	//i.Deduction()
	//p := Pay{}
	//p.Pay()
	//l := Logistics{}
	//l.SendOutGoods()

	order := NewOrder()
	order.Place()
}
