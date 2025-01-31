package main

import "fmt"

// 对象
type Obj interface {
	SendMsg(string)
	RevMsg(string)
}

// 中介者
type Mediator interface {
	SendMsg(msg string, user Obj)
}

// 用户
type User struct {
	Name     string
	mediator Mediator
}

func (u User) SendMsg(msg string) {
	fmt.Printf("%s sent: %s\n", u.Name, msg)
	u.mediator.SendMsg(msg, u)
}

func (u User) RevMsg(msg string) {
	fmt.Printf("%s rev: %s\n", u.Name, msg)
}

// 聊天室
type ChatRoom struct {
	users []User
}

func (c *ChatRoom) Register(user User) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) SendMsg(msg string, user Obj) {
	// 广播消息
	for _, u := range c.users {
		if u == user {
			continue
		}
		u.RevMsg(msg)
	}
}

func main() {
	room := ChatRoom{}
	u1 := User{Name: "zhangsan", mediator: &room}
	u2 := User{Name: "lisi", mediator: &room}
	u3 := User{Name: "wenwu", mediator: &room}

	room.Register(u1)
	room.Register(u2)
	room.Register(u3)

	u1.SendMsg("你好")
	u2.SendMsg("吃了吗")
	u3.SendMsg("我吃了")
}
