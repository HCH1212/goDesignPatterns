package main

import "fmt"

type Log struct{}

func (Log) Info(content string) {
	fmt.Println("日志落库")
}

// 代理
type ProxyLog struct {
	log *Log
}

func (p *ProxyLog) Info(content string) {
	// 延迟初始化
	if p.log == nil {
		p.log = &Log{}
	}
	// 访问前
	p.log.Info(content)
	// 访问后
	fmt.Println("记录调用了 info")
}

func main() {
	//log := ProxyLog{log: &Log{}}
	log := ProxyLog{}
	log.Info("xxxxxxxxxxxxxxx") // 没有调用 info 就不会进去初始化 &Log{}
}
