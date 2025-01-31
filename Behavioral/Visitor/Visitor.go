// 假设我们有一个文档对象模型（DOM），包含多种元素（如文本、图片、表格等）。我们需要对这些元素进行不同的操作（如导出为 PDF、导出为 HTML 等）。

package main

import "fmt"

type Element interface {
	Accept(v Visitor)
}

type Text struct {
	content string
}

// 文本操作
func (t *Text) Accept(v Visitor) {
	v.VisitText(t)
}

type Image struct {
	src string
}

func (i *Image) Accept(v Visitor) {
	v.VisitImage(i)
}

// 访问者
type Visitor interface {
	VisitImage(image *Image)
	VisitText(text *Text)
}

type PDFVisitor struct{}

func (p PDFVisitor) VisitImage(image *Image) {
	fmt.Println("pdf 获取图片的src ", image.src)
}

func (p PDFVisitor) VisitText(text *Text) {
	fmt.Println("pdf 获取文本的content ", text.content)
}

type HTMLVisitor struct{}

func (p HTMLVisitor) VisitImage(image *Image) {
	fmt.Println("html 获取图片的src ", image.src)
}

func (p HTMLVisitor) VisitText(text *Text) {
	fmt.Println("html 获取文本的content ", text.content)
}

// 定义对象结构
type Document struct {
	elements []Element
}

func (d *Document) AddElement(element Element) {
	d.elements = append(d.elements, element)
}

func (d *Document) Accept(visitor Visitor) {
	for _, element := range d.elements {
		element.Accept(visitor)
	}
}

func main() {
	docment := &Document{}
	docment.AddElement(&Image{src: "image.jpg"})
	docment.AddElement(&Text{content: "你好"})

	pdf := &PDFVisitor{}
	docment.Accept(pdf)

	html := &HTMLVisitor{}
	docment.Accept(html)
}
