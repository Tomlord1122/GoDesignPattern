/*
Prototype Pattern

This pattern allows cloning of objects without coupling to their specific classes.
It's particularly useful when the cost of creating a new object is more expensive
than copying an existing one.

Key components:
1. Prototype Interface (Prototype): Declares the cloning method
2. Concrete Prototype (Student): Implements the cloning operation
3. Client: Creates new objects by asking a prototype to clone itself

Benefits:
- Reduces subclassing needed for object creation
- Hides complexities of creating objects
- Lets you add/remove objects at runtime
- Specifies new objects by varying values
*/

package main

import "fmt"

// Prototype defines the interface for cloning objects
type Prototype interface {
	Clone() Prototype
}

// Student represents a concrete prototype that can be cloned
type Student struct {
	Name string
	Age  int
}

// Clone creates and returns a deep copy of the Student object
func (s *Student) Clone() Prototype {
	return &Student{
		Name: s.Name,
		Age:  s.Age,
	}
}

// Example usage of the Prototype Pattern
func main() {
	// Create the original prototype
	s1 := Student{
		Name: "fengfeng",
		Age:  21,
	}

	// Clone the prototype and modify its properties
	s2 := s1.Clone().(*Student)
	s2.Name = "zhangsan"
	s2.Age = 25

	// Original and cloned objects are independent
	fmt.Println(s1)
	fmt.Println(s2)
}
