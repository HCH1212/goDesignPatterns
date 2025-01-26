package main

import "fmt"

// 打印机
type Printer interface {
	PrintFile(file string) // 打印文件
}

// 创建两个打印机
type Epson struct{}

func (Epson) PrintFile(file string) {
	fmt.Println("使用爱普生打印机打印文件")
}

type Hp struct{}

func (Hp) PrintFile(file string) {
	fmt.Println("使用惠普打印机打印文件")
}

// 电脑
type Computer interface {
	Print(string)       // 打印
	SetPrinter(Printer) // 设置打印机
}

// 创建两个电脑
type Mac struct {
	printer Printer
}

func (m *Mac) Print(file string) {
	// 电脑调打印机的打印方法
	fmt.Println("使用Mac电脑")
	m.printer.PrintFile(file)
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Window struct {
	printer Printer
}

func (w *Window) Print(file string) {
	// 电脑调打印机的打印方法
	fmt.Println("使用Windows电脑")
	w.printer.PrintFile(file)
}

func (w *Window) SetPrinter(p Printer) {
	w.printer = p
}

func main() {
	w := Window{}
	//hp := Hp{}
	epson := Epson{}
	w.SetPrinter(epson)
	w.Print("12")
}
