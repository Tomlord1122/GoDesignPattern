/*
Template Method Pattern

This pattern defines the skeleton of an algorithm in a method, deferring some steps
to subclasses. It lets subclasses redefine certain steps of an algorithm without
changing the algorithm's structure. This implementation demonstrates a beverage
preparation process with variations for different drinks.

Key components:
1. Abstract Class (Template interface): Defines template method and primitive operations
2. Concrete Classes (Coffee, Tea): Implement the primitive operations
3. Template Method (MakeTemplate): Defines the algorithm skeleton
4. Hook Methods (HasAddSugar): Optional operations that subclasses can override

Benefits:
- Promotes code reuse
- Lets you enforce a skeleton algorithm
- Provides common interface for algorithm variations
- Allows fine-grained extension points
*/

package main

import "fmt"

// Template defines the interface for the beverage preparation algorithm
type Template interface {
	BoilWater()        // Step 1: Boil water
	Brew()             // Step 2: Brew the beverage
	AddSugar()         // Step 3: Add sugar (optional)
	HasAddSugar() bool // Hook method to determine if sugar should be added
}

// Coffee implements the beverage template for coffee preparation
type Coffee struct {
}

// BoilWater implements the water boiling step
func (Coffee) BoilWater() { fmt.Println("烧水") } // Boiling water

// Brew implements the coffee brewing step
func (Coffee) Brew() { fmt.Println("冲泡") } // Brewing coffee

// AddSugar implements the sugar adding step
func (Coffee) AddSugar() { fmt.Println("加糖") } // Adding sugar

// HasAddSugar specifies that coffee should have sugar
func (Coffee) HasAddSugar() bool {
	return true
}

// Tea implements the beverage template for tea preparation
type Tea struct {
}

// BoilWater implements the water boiling step
func (Tea) BoilWater() { fmt.Println("烧水") } // Boiling water

// Brew implements the tea brewing step
func (Tea) Brew() { fmt.Println("冲泡") } // Brewing tea

// AddSugar implements the sugar adding step
func (Tea) AddSugar() { fmt.Println("加糖") } // Adding sugar

// HasAddSugar specifies that tea should not have sugar
func (Tea) HasAddSugar() bool {
	return false
}

// MakeTemplate defines the template method that implements the algorithm
func MakeTemplate(tmp Template) {
	tmp.BoilWater()        // Step 1
	tmp.Brew()             // Step 2
	if tmp.HasAddSugar() { // Optional Step 3
		tmp.AddSugar()
	}
}

// Example usage of the Template Method Pattern
func main() {
	// Create beverage instances
	tea := Tea{}
	coffee := Coffee{}

	// Prepare beverages using the template method
	MakeTemplate(tea)    // Prepare tea
	MakeTemplate(coffee) // Prepare coffee
}
