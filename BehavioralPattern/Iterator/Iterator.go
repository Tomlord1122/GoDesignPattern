/*
Iterator Pattern

This pattern provides a way to access the elements of an aggregate object sequentially
without exposing its underlying representation. It encapsulates the traversal of a
collection and allows different types of traversals to be implemented.

Key components:
1. Iterator Interface: Defines interface for accessing and traversing elements
2. Concrete Iterator (BookIterator): Implements the Iterator interface
3. Aggregate Interface (implicit): Defines interface for creating Iterator
4. Concrete Aggregate (Book collection): Implements the Aggregate interface

Benefits:
- Supports variations in traversal of a collection
- Simplifies the interface of the aggregate
- Allows multiple traversals to occur simultaneously
- Provides a uniform interface for traversing different structures
*/

package main

import "fmt"

// Book represents an element in the collection
type Book struct {
	Title string // Book title
}

// BookIterator implements the iterator for a collection of books
type BookIterator struct {
	Books    []*Book // Collection of books to iterate over
	position int     // Current position in the collection
}

// NewBookIterator creates a new iterator for a collection of books
func NewBookIterator(book []*Book) *BookIterator {
	return &BookIterator{
		Books:    book,
		position: 0, // Start at the beginning of the collection
	}
}

// Iterator defines the interface for traversing the collection
type Iterator interface {
	HasNext() bool // Checks if there are more elements
	Next() *Book   // Returns the next element
}

// HasNext checks if there are more books to iterate over
func (b *BookIterator) HasNext() bool {
	if b.position >= len(b.Books) {
		return false
	}
	return true
}

// Next returns the next book in the collection
func (b *BookIterator) Next() *Book {
	if !b.HasNext() {
		return nil
	}
	book := b.Books[b.position]
	b.position++
	return book
}

// IteratorFunc demonstrates using the iterator to process all books
func IteratorFunc(iterator Iterator) {
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println(book.Title)
	}
}

// Example usage of the Iterator Pattern
func main() {
	// Create a collection of books
	book := []*Book{
		{Title: "Go开发"}, // Go Development
		{Title: "前端开发"}, // Frontend Development
		{Title: "XX开发"}, // XX Development
	}

	// Create an iterator for the book collection
	bookIterator := NewBookIterator(book)
	// Use the iterator to process all books
	IteratorFunc(bookIterator)
}
