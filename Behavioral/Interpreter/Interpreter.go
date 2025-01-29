package main

import (
	"fmt"
	"strings"
)

func main() {
	// 定义模板字符串
	const tmpl = `Hello, {{Name}}! You are {{ Age }} years old.`
	template := ParseTemplate(tmpl)
	res := template.Interpreter(&Context{
		Data: map[string]interface{}{
			"Name": "John",
			"Age":  25,
		},
	})
	fmt.Println(res)
}

// 上下文
type Context struct {
	Data map[string]any
}

// 节点
type Node interface {
	Interpreter(ctx *Context) string
}

// 文本节点
type TextNode struct {
	Content string
}

func (n *TextNode) Interpreter(ctx *Context) string {
	return n.Content
}

// 变量节点
type VarNode struct {
	Key string
}

func (n *VarNode) Interpreter(ctx *Context) string {
	val, ok := ctx.Data[n.Key]
	if !ok {
		return ""
	}
	return fmt.Sprint(val)
}

// 模板
type Template struct {
	tree []Node
}

// 解析原字符串 获取模板
func ParseTemplate(tmpl string) *Template {
	var template = new(Template)
	var index = 0
	for {
		// 查找变量表达式的起始位置
		startIndex := strings.Index(tmpl[index:], "{{")
		// 如果没有变量表达式，剩下的都是普通文本
		if startIndex == -1 {
			template.tree = append(template.tree, &TextNode{
				Content: tmpl[index:],
			})
			break
		}
		// 添加普通文本
		template.tree = append(template.tree, &TextNode{
			Content: tmpl[index : index+startIndex],
		})

		// 查找变量表达式的结束位置
		endIndex := strings.Index(tmpl[index+startIndex:], "}}")
		// 如果没有结束符，直接退出
		if endIndex == -1 {
			break
		}
		// 解析变量表达式
		key := strings.TrimSpace(tmpl[index+startIndex+2 : index+startIndex+endIndex])
		template.tree = append(template.tree, &VarNode{
			Key: key,
		})
		// 更新起始位置
		index = index + startIndex + endIndex + 2
	}
	return template
}

func (t *Template) Interpreter(ctx *Context) string {
	var s string
	for _, node := range t.tree {
		s += node.Interpreter(ctx)
	}
	return s
}
