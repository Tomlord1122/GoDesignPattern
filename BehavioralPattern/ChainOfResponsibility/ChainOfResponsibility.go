/*
Chain of Responsibility Pattern

This pattern passes requests along a chain of handlers. Upon receiving a request, each handler
decides either to process the request or to pass it to the next handler in the chain.
This implementation demonstrates a middleware chain in a web server context.

Key components:
1. Handler Interface (HandlerFun): Defines interface for handling requests
2. Concrete Handlers (AuthMiddleware, LogMiddleware): Handle requests they're responsible for
3. Client (Engine): Initiates requests to the chain of handlers
4. Context: Holds request information and controls chain traversal

Benefits:
- Decouples senders and receivers
- Adds flexibility in assigning responsibilities to objects
- Allows dynamic changes to the chain
- Promotes single responsibility principle
*/

package main

import (
	"fmt"
	"net/http"
)

// Context holds request information and controls middleware chain execution
type Context struct {
	request  *http.Request       // The HTTP request being processed
	w        http.ResponseWriter // The HTTP response writer
	index    int                 // Current position in the middleware chain
	handlers []HandlerFun        // Slice of middleware handlers
}

// Next advances to the next middleware in the chain
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

// Abort stops the middleware chain execution
func (c *Context) Abort() {
	c.index = len(c.handlers)
}

// HandlerFun defines the middleware handler function type
type HandlerFun func(*Context)

// Engine represents the web server that manages the middleware chain
type Engine struct {
	handlers []HandlerFun // Slice of middleware handlers
}

// Use adds a middleware to the chain
func (e *Engine) Use(f HandlerFun) {
	e.handlers = append(e.handlers, f)
}

// ServeHTTP implements the http.Handler interface and initiates the middleware chain
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := &Context{
		request:  r,
		w:        w,
		index:    -1,
		handlers: e.handlers,
	}
	context.Next()
}

// AuthMiddleware handles authentication in the middleware chain
func AuthMiddleware(c *Context) {
	fmt.Println("认证中间件") // Authentication middleware
}

// LogMiddleware handles logging in the middleware chain
func LogMiddleware(c *Context) {
	fmt.Println("日志中间件") // Logging middleware
	c.Next()             // Continue to next middleware
}

// Example usage of the Chain of Responsibility Pattern
func main() {
	// Create the engine (client)
	r := &Engine{}

	// Add middleware handlers to the chain
	r.Use(LogMiddleware)  // First handler in chain
	r.Use(AuthMiddleware) // Second handler in chain

	// Start the web server
	fmt.Println("web server on :8080")
	http.ListenAndServe(":8080", r)
}
