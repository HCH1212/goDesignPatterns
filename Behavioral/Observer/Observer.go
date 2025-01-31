package main

import "fmt"

// 观察者
type Observer interface {
	RevMsg(wd int) // 接受温度变化
}

type User struct {
	Name string
}

func (u User) RevMsg(wd int) {
	fmt.Printf("%s现在温度%d\n", u.Name, wd)
}

// 被观察者
type Subject interface {
	SendMsg(wd int)
	Notify()
	RegisterObserver(Observer)
}

type WeatherZhan struct {
	observerList []Observer
	wd           int
}

func (w *WeatherZhan) SendMsg(wd int) {
	w.wd = wd
	w.Notify()
}

func (w *WeatherZhan) Notify() {
	for _, observer := range w.observerList {
		observer.RevMsg(w.wd)
	}
}

func (w *WeatherZhan) RegisterObserver(o Observer) {
	w.observerList = append(w.observerList, o)
}

func main() {
	zhan := &WeatherZhan{}
	u1 := User{Name: "zhangsan"}
	u2 := User{Name: "lisi"}

	zhan.RegisterObserver(u1)
	zhan.RegisterObserver(u2)

	zhan.SendMsg(8)
	zhan.SendMsg(0)
}
