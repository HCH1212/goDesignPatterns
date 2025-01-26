package main

import "fmt"

// 组件接口：文件系统节点
type Node interface {
	Display(ident string)
}

// 文件
type File struct {
	Name string
}

func (f *File) Display(ident string) {
	fmt.Println(ident + f.Name)
}

// 文件夹
type Dir struct {
	Name     string
	Children []Node
}

func (d *Dir) Display(ident string) {
	fmt.Println(ident + d.Name)
	for _, child := range d.Children {
		child.Display(ident + "  ")
	}
}

func main() {
	root := Dir{
		Name: "Creational",
		Children: []Node{
			&Dir{
				Name: "AbstractFactory",
				Children: []Node{
					&File{
						Name: "AbstractFactory.go",
					},
				},
			},
			&File{
				Name: "main.go",
			},
		},
	}
	root.Display("")
}
