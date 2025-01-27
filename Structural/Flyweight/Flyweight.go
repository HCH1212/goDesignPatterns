// 各种池

package main

import "fmt"

// 一个数据库连接
type DB struct {
	id int
}

func (db DB) Query(string2 string) {
	fmt.Printf("使用 %d 连接对象 进行查询操作 %s\n", db.id, string2)
}

func NewDB(id int) *DB {
	return &DB{id}
}

// 连接池
type DBPool struct {
	pool   map[int]*DB
	nextID int
}

func NewDBPool(num int) *DBPool {
	pool := make(map[int]*DB, num) // 真实场景需要考虑并发条件，加锁
	for i := 0; i < num; i++ {
		pool[i] = NewDB(i)
	}
	return &DBPool{pool, num - 1}
}

// 从连接池获取DB
func (p *DBPool) GetDB() *DB {
	if len(p.pool) > 0 {
		for id, db := range p.pool {
			delete(p.pool, id)
			return db
		}
	}
	// 没有就创建
	p.nextID++
	db := NewDB(p.nextID)
	return db
}

// 放回到连接池
func (p *DBPool) Release(db *DB) {
	p.pool[db.id] = db
}

func main() {
	pool := NewDBPool(1)

	db1 := pool.GetDB()
	db1.Query("select * from xxx")

	db2 := pool.GetDB()
	db2.Query("select * from xxx")
	pool.Release(db2)

	db3 := pool.GetDB()
	db3.Query("select * from xxx")
	pool.Release(db3)
}
