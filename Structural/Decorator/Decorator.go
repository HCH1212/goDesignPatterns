//假设你正在开发一个 Web 服务，其中有一个核心功能是处理用户请求。现在，你需要在不修改核心功能代码的情况下，为请求处理添加以下功能：
//
//1. 日志记录：记录每个请求的详细信息。
//2. 性能监控：记录每个请求的处理时间。

package main

import (
	"fmt"
	"time"
)

// 用户请求
type ReqI interface {
	Handler(url string) string
}

type Req struct{}

func (r Req) Handler(url string) string {
	fmt.Println("请求 " + url)
	time.Sleep(100 * time.Millisecond)
	return ""
}

// 日志装饰器
type LogReqDecorator struct {
	req ReqI
}

func (l LogReqDecorator) Handler(url string) string {
	fmt.Println("日志记录前")
	res := l.req.Handler(url)
	fmt.Println("日志记录后")
	return res
}

// 性能监控装饰器
type MonitorDecorator struct {
	req ReqI
}

func (m MonitorDecorator) Handler(url string) string {
	t1 := time.Now()
	res := m.req.Handler(url)
	sub := time.Since(t1)
	fmt.Println("耗时：", sub)
	return res
}

func main() {
	req := Req{}
	req1 := LogReqDecorator{req: req}
	req2 := MonitorDecorator{req: req1} // 嵌套进去，实现两个功能
	req2.Handler("http://www.baidu.com")
}
