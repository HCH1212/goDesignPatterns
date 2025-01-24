package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}
type Student struct {
	Name string
	Age  int
}

func (s *Student) Clone() Prototype {
	return &Student{
		Name: s.Name,
		Age:  s.Age,
	}
}

func main() {
	s1 := Student{
		Name: "s1",
		Age:  20,
	}
	s2 := s1.Clone().(*Student)
	s2.Name = "s2"
	s2.Age = 30

	fmt.Println(s1)
	fmt.Println(s2)
}
