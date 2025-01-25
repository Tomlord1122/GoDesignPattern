/*
Visitor Pattern

This pattern represents an operation to be performed on elements of an object structure.
It lets you define a new operation without changing the classes of the elements on which
it operates. This implementation demonstrates a document object model (DOM) where different
visitors can process elements in different ways (PDF export, HTML export, etc.).

Key components:
1. Visitor Interface: Declares visit operations for each element type
2. Concrete Visitors (PDFVisitor, HTMLVisitor): Implement operations for each element
3. Element Interface: Declares accept operation that takes visitor as argument
4. Concrete Elements (Text, Image): Implement accept operation
5. Object Structure (Document): Can enumerate its elements

Benefits:
- Makes adding new operations easy
- Gathers related operations and separates unrelated ones
- Combines operations that work on different class hierarchies
- Accumulates state as visitor traverses structure
*/

package main

import "fmt"

// Element defines the interface for objects that can be visited
type Element interface {
	Accept(v Visitor) // Accept a visitor to perform an operation
}

// Text represents a text element in the document
type Text struct {
	content string // Text content
}

// Accept implements the visitor acceptance for text elements
func (t *Text) Accept(v Visitor) {
	v.VisitText(t)
}

// Image represents an image element in the document
type Image struct {
	src string // Image source path
}

// Accept implements the visitor acceptance for image elements
func (t *Image) Accept(v Visitor) {
	v.VisitImage(t)
}

// Visitor defines the interface for different document processors
type Visitor interface {
	VisitText(text *Text)    // Process a text element
	VisitImage(image *Image) // Process an image element
}

// PDFVisitor implements PDF export functionality
type PDFVisitor struct {
}

// VisitText processes text elements for PDF export
func (p PDFVisitor) VisitText(text *Text) {
	fmt.Println("pdf 文本内容", text.content) // PDF text content
}

// VisitImage processes image elements for PDF export
func (p PDFVisitor) VisitImage(image *Image) {
	fmt.Println("pdf 获取图片的src", image.src) // PDF get image source
}

// HTMLVisitor implements HTML export functionality
type HTMLVisitor struct {
}

// VisitText processes text elements for HTML export
func (p HTMLVisitor) VisitText(text *Text) {
	fmt.Println("html 文本内容", text.content) // HTML text content
}

// VisitImage processes image elements for HTML export
func (p HTMLVisitor) VisitImage(image *Image) {
	fmt.Println("html 获取图片的src", image.src) // HTML get image source
}

// Document represents the object structure containing elements
type Document struct {
	elements []Element // Collection of document elements
}

// AddElement adds a new element to the document
func (d *Document) AddElement(element Element) {
	d.elements = append(d.elements, element)
}

// Accept implements the visitor pattern for the entire document
func (d *Document) Accept(visitor Visitor) {
	for _, element := range d.elements {
		element.Accept(visitor)
	}
}

// Example usage of the Visitor Pattern
func main() {
	// Create document structure
	document := &Document{}

	// Add various elements
	document.AddElement(&Text{content: "你好"})   // Hello
	document.AddElement(&Text{content: "xxx"})  // xxx
	document.AddElement(&Image{src: "abc.png"}) // Image source

	// Process document using PDF visitor
	pdf := &PDFVisitor{}
	document.Accept(pdf)

	// Process document using HTML visitor
	html := &HTMLVisitor{}
	document.Accept(html)
}
