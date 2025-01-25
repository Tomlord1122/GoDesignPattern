/*
Composite Pattern

This pattern lets you compose objects into tree structures to represent part-whole hierarchies.
It lets clients treat individual objects and compositions of objects uniformly.
This implementation demonstrates a file system structure where both files and directories
are treated as nodes in the tree.

Key components:
1. Component (Node): Interface that defines common operations for both leaf and composite objects
2. Leaf (File): Represents end objects of the composition with no children
3. Composite (Dir): Represents objects that have children and operations to manipulate them

Benefits:
- Makes it easy to add new kinds of components
- Provides flexibility in building tree structures
- Clients can treat composite and individual objects uniformly
- Simplifies client code by allowing uniform treatment of objects
*/

package main

import "fmt"

// Node defines the common interface for both File and Dir objects
type Node interface {
	Display(ident string) // Display the node with proper indentation
}

// File represents a leaf node in the file system tree
type File struct {
	Name string
}

// Display shows the file name with the given indentation
func (n File) Display(ident string) {
	fmt.Println(ident + n.Name)
}

// Dir represents a composite node that can contain other nodes (files or directories)
type Dir struct {
	Name     string // Directory name
	children []Node // List of child nodes (files or subdirectories)
}

// Display shows the directory structure recursively
func (n Dir) Display(ident string) {
	// Display current directory name
	fmt.Println(ident + n.Name)
	// Recursively display all children with increased indentation
	for _, child := range n.children {
		child.Display(ident + "  ")
	}
}

// Example usage of the Composite Pattern
func main() {
	// Create a sample directory structure
	root := Dir{
		Name: "CreationalPatterns",
		children: []Node{
			Dir{
				Name: "AbstractFactory",
				children: []Node{
					File{
						Name: "AbstractFactory.go",
					},
				},
			},
			File{
				Name: "main.go",
			},
		},
	}
	// Display the entire directory structure
	root.Display("")
}
