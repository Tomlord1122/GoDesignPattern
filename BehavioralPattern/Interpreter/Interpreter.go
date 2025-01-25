/*
Interpreter Pattern

This pattern defines a grammar for interpreting the sentences in a simple language
and provides an interpreter to handle them. This implementation demonstrates a simple
template engine that can interpret and replace variables in a template string.

Key components:
1. Abstract Expression (Node): Declares an interface for executing an operation
2. Terminal Expression (TextNode): Implements interpretation for terminal symbols
3. Non-terminal Expression (VarNode): Implements interpretation for non-terminal symbols
4. Context: Contains information global to the interpreter
5. Client: Builds the abstract syntax tree and invokes the interpretation

Benefits:
- Makes it easy to change and extend the grammar
- Implements the grammar rules as classes
- Separates grammar from the interpretation
- Adds new ways to interpret expressions
*/

package main

import (
	"fmt"
	"strings"
)

// Example usage of the Interpreter Pattern
func main() {
	// Define the template string with variables to interpret
	const tmpl = `Hello, {{ Name }}! You are {{Age}} years old.`
	// Parse the template into an abstract syntax tree
	template := ParseTemplate(tmpl)
	// Interpret the template with provided context
	res := template.Interpreter(&Context{
		Data: map[string]any{
			"Name": "fengfeng",
			"Age":  21,
		},
	})
	fmt.Println(res)
}

// Context holds the variables for interpretation
type Context struct {
	Data map[string]any // Map of variable names to their values
}

// Node defines the interface for all expression nodes
type Node interface {
	Interpreter(*Context) string
}

// TextNode represents literal text in the template
type TextNode struct {
	Content string // The literal text content
}

// Interpreter returns the literal text content
func (t *TextNode) Interpreter(ctx *Context) string {
	return t.Content
}

// VarNode represents a variable in the template
type VarNode struct {
	Key string // The variable name
}

// Interpreter looks up and returns the variable value from the context
func (t *VarNode) Interpreter(ctx *Context) string {
	val, ok := ctx.Data[t.Key]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v", val)
}

// Template represents the parsed template with its abstract syntax tree
type Template struct {
	tree []Node // List of nodes in the template
}

// ParseTemplate parses a template string into an abstract syntax tree
func ParseTemplate(tmpl string) *Template {
	var template = new(Template)
	var index = 0
	for {
		// Find the next variable start marker
		startIndex := strings.Index(tmpl[index:], "{{")
		if startIndex == -1 {
			// No more variables, add remaining text as TextNode
			template.tree = append(template.tree, &TextNode{
				Content: tmpl[index:],
			})
			break
		}
		// Add text before the variable as TextNode
		template.tree = append(template.tree, &TextNode{
			Content: tmpl[index : index+startIndex],
		})
		// Find the variable end marker
		endIndex := strings.Index(tmpl[index+startIndex:], "}}")
		if endIndex == -1 {
			break
		}
		// Extract and add the variable name as VarNode
		key := strings.TrimSpace(tmpl[index+startIndex+2 : index+startIndex+endIndex])
		template.tree = append(template.tree, &VarNode{
			Key: key,
		})
		index = index + startIndex + endIndex + 2
	}
	return template
}

// Interpreter processes the template by interpreting all nodes
func (t *Template) Interpreter(ctx *Context) string {
	var s string
	// Interpret each node in sequence
	for _, node := range t.tree {
		s += node.Interpreter(ctx)
	}
	return s
}
