package main

import "fmt"

type Book struct {
	Title string
}

// 书本集合
type BookIterator struct {
	Books    []*Book
	position int
}

func NewBookIterator(books []*Book) *BookIterator {
	return &BookIterator{
		Books:    books,
		position: 0,
	}
}

// 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() *Book
}

func (iter *BookIterator) HasNext() bool {
	return iter.position < len(iter.Books)
}

func (iter *BookIterator) Next() *Book {
	if iter.HasNext() {
		book := iter.Books[iter.position]
		iter.position++
		return book
	}
	return nil
}

func IteratorFunc(iterator Iterator) {
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println(book.Title)
	}
}

func main() {
	book := []*Book{
		{"GO"},
		{"JAVA"},
		{"C++"},
	}
	bookIterator := NewBookIterator(book)
	IteratorFunc(bookIterator)
}
