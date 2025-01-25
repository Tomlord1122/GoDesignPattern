/*
Builder Pattern

This pattern separates the construction of a complex object from its representation,
allowing the same construction process to create different representations.

Key components:
1. Product (House): The complex object being built
2. Builder Interface (HouseBuilder): Specifies abstract interface for creating parts of the product
3. Concrete Builder (Bao): Constructs and assembles parts of the product
4. Director (Boos): Constructs the object using the Builder interface

Benefits:
- Allows fine control over the construction process
- Isolates code for construction and representation
- Provides finer control over the construction process
- Allows objects to be constructed step by step
*/

package main

import "fmt"

// House represents the complex object being built
type House struct {
	Door   string
	Window string
}

// HouseBuilder specifies the interface for building parts of a house
type HouseBuilder interface {
	BuildDoor(val string)
	BuildWindow(val string)
	GetHouse() *House
}

// Bao is a concrete builder that implements the HouseBuilder interface
type Bao struct {
	house *House
}

// NewBao creates a new builder instance with an empty house
func NewBao() *Bao {
	return &Bao{
		house: &House{},
	}
}

// BuildDoor builds the door part of the house
func (b *Bao) BuildDoor(val string) {
	b.house.Door = val
	fmt.Println("门建造成功") // Door construction successful
}

// BuildWindow builds the window part of the house
func (b *Bao) BuildWindow(val string) {
	b.house.Window = val
	fmt.Println("窗户建造成功") // Window construction successful
}

// GetHouse returns the constructed house
func (b *Bao) GetHouse() *House {
	return b.house
}

// Boos acts as the director in the builder pattern
type Boos struct {
	builder HouseBuilder
}

// NewBoos creates a new director with the specified builder
func NewBoos(bao *Bao) *Boos {
	return &Boos{
		builder: bao,
	}
}

// GetHouse directs the construction process and returns the final product
func (b Boos) GetHouse() *House {
	// Define the construction steps sequence
	b.builder.BuildDoor("门")   // Build door
	b.builder.BuildWindow("窗") // Build window
	return b.builder.GetHouse()
}

// Example usage of the Builder Pattern
func main() {
	// Create the builder
	b := NewBao()
	// Create the director with the builder
	boos := NewBoos(b)
	// Get the constructed house
	fmt.Println(boos.GetHouse())
}
