package main

import "fmt"

// 命令
type Command interface {
	Execute()
}

// 打印命令
type PrintCommand struct {
	Content string
}

func (p PrintCommand) Execute() {
	fmt.Println("打印消息: " + p.Content)
}

// 发邮件命令
type SendEmail struct {
	To      string
	Content string
}

func (s SendEmail) Execute() {
	fmt.Println("发送邮件: "+s.To, s.Content)
}

// 一个任务队列
type TaskQueue struct {
	Queue []Command
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{}
}

func (t *TaskQueue) AddCommand(command Command) {
	t.Queue = append(t.Queue, command)
}

func (t *TaskQueue) Command() {
	for _, command := range t.Queue {
		command.Execute()
	}
}

func main() {
	queue := NewTaskQueue()
	queue.AddCommand(PrintCommand{
		Content: "你好",
	})
	queue.AddCommand(SendEmail{
		To:      "abc@qq.com",
		Content: "i love you",
	})
	queue.Command()
}
