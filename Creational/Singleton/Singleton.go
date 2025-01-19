package main

import (
	"fmt"
	"sync"
)

type DB struct {
	Dsn string
}

var db *DB // 外部使用只用调用GetDB()

func InitDB(dsn string) *DB {
	return &DB{dsn}
}

var once sync.Once

func GetDB() *DB {
	once.Do(func() {
		db = InitDB("127.0.0.1:3306")
	})
	return db
}

func main() {
	d1 := GetDB()
	d2 := GetDB()
	d3 := GetDB()

	fmt.Printf("%p\n", d1)
	fmt.Printf("%p\n", d2)
	fmt.Printf("%p\n", d3)
}
