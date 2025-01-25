/*
Memento Pattern

This pattern lets you save and restore the previous state of an object without
revealing the details of its implementation. It captures an object's internal
state and externally saves it so that the object can be restored to this state later.

Key components:
1. Originator (TextNode): Creates and stores states in mementos
2. Memento Interface: Provides interface for accessing state
3. Concrete Memento (TextMemento): Stores internal state of the Originator
4. Caretaker (Manage): Keeps track of multiple mementos but never modifies them

Benefits:
- Provides a way to recover from errors
- Simplifies the originator
- Provides easy-to-implement recovery capability
- Encapsulates the internal state management
*/

package main

import "fmt"

// Node defines the interface for objects that can have their state managed
type Node interface {
	SetState(state string) // Set the object's state
	GetState() string      // Get the object's current state
}

// TextNode represents an object whose state can be saved and restored
type TextNode struct {
	state string // Current state of the text node
}

// SetState updates the text node's state
func (t *TextNode) SetState(state string) {
	t.state = state
}

// GetState retrieves the text node's current state
func (t *TextNode) GetState() string {
	return t.state
}

// Save creates a memento containing the current state
func (t *TextNode) Save() Memento {
	return &TextMemento{
		state: t.state,
	}
}

// Memento defines the interface for accessing saved states
type Memento interface {
	GetState() string // Get the saved state
}

// TextMemento stores the state of a TextNode
type TextMemento struct {
	state string // Saved state
}

// GetState retrieves the saved state from the memento
func (t TextMemento) GetState() string {
	return t.state
}

// Manage acts as the caretaker for mementos
type Manage struct {
	states []Memento // Collection of saved states
}

// Save stores a memento
func (m *Manage) Save(t Memento) {
	m.states = append(m.states, t)
}

// Back retrieves a memento at a specific index
func (m *Manage) Back(index int) Memento {
	return m.states[index]
}

// Example usage of the Memento Pattern
func main() {
	// Create the caretaker
	manage := Manage{}

	// Create the originator
	text := TextNode{}

	// Set and save first state
	text.SetState("1")
	manage.Save(text.Save())

	// Set and save second state
	text.SetState("2")
	manage.Save(text.Save())

	// Show current state
	fmt.Println(text.GetState())
	// Show restored state from first save
	fmt.Println(manage.Back(0).GetState())
}
