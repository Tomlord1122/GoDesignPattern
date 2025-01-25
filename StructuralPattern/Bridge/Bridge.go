/*
Bridge Pattern

This pattern decouples an abstraction from its implementation so that the two can vary independently.
It's particularly useful when both the abstraction and its implementation need to be extended
using inheritance.

Key components:
1. Abstraction (Computer): Defines the abstract interface and maintains reference to implementor
2. Refined Abstraction (Mac, Windows): Extends the abstraction
3. Implementor (Printer): Defines the interface for implementation classes
4. Concrete Implementor (Epson, Hp): Implements the Implementor interface

Benefits:
- Decouples interface from implementation
- Improves extensibility
- Hides implementation details from clients
- Allows sharing of implementations among objects
*/

package main

import "fmt"

// Printer defines the interface for implementation classes (Implementor)
type Printer interface {
	PrintFile(file string) // Print a file
}

// Epson implements the Printer interface for Epson printers
type Epson struct {
}

// PrintFile implements the printing functionality for Epson printers
func (Epson) PrintFile(file string) {
	fmt.Println("使用爱普生打印机打印文件") // Using Epson printer to print file
}

// Hp implements the Printer interface for HP printers
type Hp struct {
}

// PrintFile implements the printing functionality for HP printers
func (Hp) PrintFile(file string) {
	fmt.Println("使用惠普打印机打印文件") // Using HP printer to print file
}

// Computer defines the abstraction interface
type Computer interface {
	Print(string)       // Print operation
	SetPrinter(Printer) // Set the printer to use
}

// Mac represents a refined abstraction for Mac computers
type Mac struct {
	printer Printer // Reference to the printer implementation
}

// Print implements the printing operation for Mac computers
func (m *Mac) Print(file string) {
	fmt.Println("使用mac电脑") // Using Mac computer
	m.printer.PrintFile(file)
}

// SetPrinter sets the printer implementation for Mac
func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

// Windows represents a refined abstraction for Windows computers
type Windows struct {
	printer Printer // Reference to the printer implementation
}

// Print implements the printing operation for Windows computers
func (m *Windows) Print(file string) {
	fmt.Println("使用windows电脑") // Using Windows computer
	m.printer.PrintFile(file)
}

// SetPrinter sets the printer implementation for Windows
func (m *Windows) SetPrinter(printer Printer) {
	m.printer = printer
}

// Example usage of the Bridge Pattern
func main() {
	// Create a Windows computer
	w := Windows{}
	// Create an Epson printer
	ep := Epson{}

	// Connect the Windows computer with the Epson printer
	w.SetPrinter(ep)
	// Print using the Windows-Epson bridge
	w.Print("xx")
}
