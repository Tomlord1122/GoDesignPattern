/*
Facade Pattern

This pattern provides a unified interface to a set of interfaces in a subsystem.
It defines a higher-level interface that makes the subsystem easier to use by hiding its complexity.
This implementation demonstrates an order processing system that coordinates inventory,
payment, and logistics operations.

Key components:
1. Facade (Order): Provides a simplified interface to complex subsystem
2. Subsystem Classes (Inventory, Pay, Logistics): Implement subsystem functionality
3. Client: Uses the facade to interact with the subsystem

Benefits:
- Simplifies client interaction with complex subsystems
- Promotes loose coupling between subsystems and clients
- Makes subsystems easier to use and reduces dependencies
- Provides a context-specific interface to more general functionality
*/

package main

import "fmt"

// Inventory manages product inventory operations
type Inventory struct {
}

// Deduction reduces the inventory count
func (k Inventory) Deduction() {
	fmt.Println("扣库存") // Deduct from inventory
}

// Pay handles payment processing
type Pay struct {
}

// Pay processes the payment
func (k Pay) Pay() {
	fmt.Println("支付") // Process payment
}

// Logistics manages shipping and delivery
type Logistics struct {
}

// SendOutGoods initiates the shipping process
func (k Logistics) SendOutGoods() {
	fmt.Println("发货") // Ship the goods
}

// Order serves as the facade, coordinating all order processing operations
type Order struct {
	inventory *Inventory
	pay       *Pay
	logistics *Logistics
}

// NewOrder creates a new Order facade with all required subsystems
func NewOrder() *Order {
	return &Order{
		inventory: &Inventory{},
		pay:       &Pay{},
		logistics: &Logistics{},
	}
}

// Place processes the order by coordinating all necessary operations
func (o Order) Place() {
	// Execute operations in the required sequence
	o.inventory.Deduction()    // First deduct from inventory
	o.pay.Pay()                // Then process payment
	o.logistics.SendOutGoods() // Finally ship the goods
}

// Example usage of the Facade Pattern
func main() {
	// Create and use the facade
	o := NewOrder()
	// Process the order with a single method call
	o.Place()
}
