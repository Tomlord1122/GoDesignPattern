/*
Flyweight Pattern

This pattern minimizes memory usage by sharing as much data as possible with similar objects.
It's particularly useful when dealing with a large number of similar objects.
This implementation demonstrates a database connection pool where connection objects
are reused instead of creating new ones for each request.

Key components:
1. Flyweight (DB): Contains the shared state that can be reused
2. Flyweight Factory (DBPool): Creates and manages flyweight objects
3. Context: The extrinsic state that makes each instance unique
4. Client: Uses flyweights through the factory

Benefits:
- Reduces memory usage when dealing with many similar objects
- Improves performance through object reuse
- Centralizes state management for shared objects
- Supports efficient scaling of applications
*/

package main

import "fmt"

// DB represents a database connection (the Flyweight)
type DB struct {
	id int // Intrinsic state that can be shared
}

// NewDB creates a new database connection with a specific ID
func NewDB(id int) *DB {
	return &DB{id: id}
}

// Query performs a database query using the connection
func (db DB) Query(string2 string) {
	fmt.Printf("使用 %d 连接对象 进行查询操作 %s\n", db.id, string2) // Using connection object [id] to query [string]
}

// DBPool manages a pool of database connections (Flyweight Factory)
type DBPool struct {
	pool   map[int]*DB // Map of available DB connections
	nextID int         // Next available connection ID
}

// NewDBPool creates a new connection pool with initial connections
func NewDBPool(num int) *DBPool {
	var pool = map[int]*DB{}
	// Initialize the pool with the specified number of connections
	for i := 0; i < num; i++ {
		pool[i] = NewDB(i)
	}
	return &DBPool{
		pool:   pool,
		nextID: num - 1,
	}
}

// GetDB retrieves a database connection from the pool
func (p *DBPool) GetDB() *DB {
	// Try to reuse an existing connection from the pool
	if len(p.pool) > 0 {
		for id, db := range p.pool {
			delete(p.pool, id)
			return db
		}
	}
	// If no connections are available, create a new one
	p.nextID++
	db := NewDB(p.nextID)
	return db
}

// Release returns a database connection back to the pool
func (p *DBPool) Release(db *DB) {
	p.pool[db.id] = db
}

// Example usage of the Flyweight Pattern
func main() {
	// Create a connection pool with 1 initial connection
	pool := NewDBPool(1)

	// Use and release the same connection multiple times
	db1 := pool.GetDB()
	db1.Query("select * from xxx")
	pool.Release(db1)

	db2 := pool.GetDB()
	db2.Query("select * from xxx")
	pool.Release(db2)

	db3 := pool.GetDB()
	db3.Query("select * from xxx")
	pool.Release(db3)
}
