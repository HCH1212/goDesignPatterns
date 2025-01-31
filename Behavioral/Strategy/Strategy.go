// 假设我们有一个支付系统，支持多种支付方式（如信用卡支付、支付宝支付、微信支付等）。我们可以使用策略模式来实现支付方式的选择。

package main

import "fmt"

type Pay interface {
	Pay(money int64)
}

type AliPay struct{}

func (AliPay) Pay(money int64) {
	fmt.Println("使用支付宝支付", money)
}

type WeiPay struct{}

func (WeiPay) Pay(money int64) {
	fmt.Println("使用微信支付", money)
}

// 支付策略
type PayStrategy struct {
	pay Pay
}

func (p *PayStrategy) SetPay(pay Pay) {
	p.pay = pay
}

func (p *PayStrategy) Pay(money int64) {
	p.pay.Pay(money)
}

func main() {
	aliPay := &AliPay{}
	weiPay := &WeiPay{}
	payStrategy := &PayStrategy{}

	payStrategy.SetPay(aliPay)
	payStrategy.Pay(12)

	payStrategy.SetPay(weiPay)
	payStrategy.Pay(19)
}
