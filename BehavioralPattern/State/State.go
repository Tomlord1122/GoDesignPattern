/*
State Pattern

This pattern allows an object to alter its behavior when its internal state changes.
The object will appear to change its class. This implementation demonstrates a
simple switch that can toggle between on and off states.

Key components:
1. State Interface: Defines interface for state-specific behavior
2. Concrete States (OnState, OffState): Implements behavior for specific states
3. Context: Maintains reference to current state and delegates state-specific behavior
4. Client: Uses the context to change states

Benefits:
- Localizes state-specific behavior
- Makes state transitions explicit
- Eliminates large state machine conditionals
- Allows states to be shared among contexts
*/

package main

import "fmt"

// State defines the interface for different states
type State interface {
	Switch(context *Context) // Handle state transition
}

// Context maintains the current state and delegates behavior to it
type Context struct {
	state State // Current state
}

// SetState changes the current state
func (c *Context) SetState(state State) {
	c.state = state
}

// Switch delegates the switch operation to the current state
func (c *Context) Switch() {
	c.state.Switch(c)
}

// OnState represents the "on" state of the switch
type OnState struct {
}

// Switch handles the transition from "on" to "off" state
func (OnState) Switch(context *Context) {
	fmt.Println("开关关闭") // Switch turned off
	context.SetState(&OffState{})
}

// OffState represents the "off" state of the switch
type OffState struct {
}

// Switch handles the transition from "off" to "on" state
func (OffState) Switch(context *Context) {
	fmt.Println("开关打开") // Switch turned on
	context.SetState(&OnState{})
}

// Example usage of the State Pattern
func main() {
	// Create context with initial "off" state
	c := &Context{
		state: OffState{},
	}

	// Toggle the switch multiple times
	c.Switch() // Turn on
	c.Switch() // Turn off
	c.Switch() // Turn on
}
