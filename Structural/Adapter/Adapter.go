package main

import "fmt"

// 自己的支付模式
type AliPay struct{}

func (a AliPay) GetPayPage(price int) string {
	return "支付宝支付链接"
}

// 支付宝适配微信
type AliPayAdapter struct {
	aliPay *AliPay
}

func (a AliPayAdapter) PayPage(price int) string {
	return a.aliPay.GetPayPage(price)
}

// 对方的支付模式
type WeiXinPay struct{}

func (w WeiXinPay) PayPage(price int) string {
	return "微信支付链接"
}

// 接口方法
type PayI interface {
	PayPage(price int) string
}

func PayPage(pi PayI, price int) string {
	return pi.PayPage(price)
}

func main() {
	fmt.Println(PayPage(WeiXinPay{}, 1))
	fmt.Println(PayPage(AliPayAdapter{aliPay: &AliPay{}}, 1))
}
