// 类似于中间件中的Next和Abort

package main

import (
	"fmt"
	"net/http"
)

type Context struct {
	request  *http.Request
	w        http.ResponseWriter
	index    int
	handlers []HandlerFunc
}

func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c) // 执行
	}
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}

type HandlerFunc func(*Context)

// 引擎
type Engine struct {
	handlers []HandlerFunc
}

func (e *Engine) Use(f HandlerFunc) {
	e.handlers = append(e.handlers, f)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := &Context{r, w, -1, e.handlers}
	context.Next()
}

func AuthMiddleware(c *Context) {
	fmt.Println("认证中间件")
}

func LogMiddleware(c *Context) {
	fmt.Println("日志中间件")
	c.Next()
}

func main() {
	r := &Engine{}
	r.Use(LogMiddleware)
	r.Use(AuthMiddleware)

	fmt.Println("web server on :8080")
	http.ListenAndServe(":8080", r)
}
