package main

import "fmt"

// 支付产品
type Pay interface {
	PayPage(price int64) string
}
type AliPay struct {
}

func (AliPay) PayPage(price int64) string {
	return "支付宝支付"
}

type WeiXinPay struct {
}

func (WeiXinPay) PayPage(price int64) string {
	return "微信支付"
}

// 退款产品
type ReFund interface {
	ReFund(no string) error
}

type AliReFund struct {
}

func (AliReFund) ReFund(no string) error {
	fmt.Println("支付宝退款")
	return nil
}

type WeiXinReFund struct{}

func (WeiXinReFund) ReFund(no string) error {
	fmt.Println("微信退款")
	return nil
}

// 创建工厂接口
type PayFactory interface {
	CreatePay() Pay
	CreateReFund() ReFund
}

// 创建两个厂
type AliPayFactory struct {
}

func (AliPayFactory) CreatePay() Pay {
	// 专门管支付宝支付的相关逻辑
	return AliPay{}
}
func (AliPayFactory) CreateReFund() ReFund {
	return AliReFund{}
}

type WeiXinPayFactory struct {
}

func (WeiXinPayFactory) CreatePay() Pay {
	return WeiXinPay{}
}
func (WeiXinPayFactory) CreateReFund() ReFund {
	return WeiXinReFund{}
}

func main() {
	aliPayFactory := AliPayFactory{}
	aliPay := aliPayFactory.CreatePay()
	fmt.Println(aliPay.PayPage(15))
	_ = aliPayFactory.CreateReFund().ReFund("1")

	weiXinPayFactory := WeiXinPayFactory{}
	weiXinPay := weiXinPayFactory.CreatePay()
	fmt.Println(weiXinPay.PayPage(15))
	_ = weiXinPayFactory.CreateReFund().ReFund("1")
}
