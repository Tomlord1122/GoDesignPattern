/*
Singleton Pattern

This pattern ensures a class has only one instance and provides a global point of access to it.
It's particularly useful for managing shared resources like database connections.

Key components:
1. Private constructor (implicitly private in Go)
2. Private instance variable (db *DB)
3. Public access method (GetDB())
4. Thread-safe initialization using sync.Once

Benefits:
- Ensures only one instance exists
- Provides global access point to that instance
- Lazy initialization
- Thread-safe implementation
*/

package main

import (
	"fmt"
	"sync"
)

// DB represents a database connection
type DB struct {
}

// db is the private instance variable that holds the singleton instance
var db *DB

// InitDB creates a new database connection
// In a real application, this would handle the actual database connection setup
func InitDB(dsn string) *DB {
	return &DB{}
}

// once ensures thread-safe initialization of the singleton
var once sync.Once

// GetDB provides global access to the singleton database instance
// It ensures the instance is created only once, even in concurrent scenarios
func GetDB() *DB {
	once.Do(func() {
		db = InitDB("")
	})
	return db
}

// Example usage of the Singleton Pattern
func main() {
	// Get multiple references to the DB instance
	d1 := GetDB()
	d2 := GetDB()
	d3 := GetDB()

	// Print memory addresses to prove they're the same instance
	fmt.Printf("%p\n", d1)
	fmt.Printf("%p\n", d2)
	fmt.Printf("%p\n", d3)
}
