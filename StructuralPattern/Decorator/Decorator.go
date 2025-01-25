/*
Decorator Pattern

This pattern allows behavior to be added to individual objects dynamically by placing
these objects inside wrapper objects that contain the behaviors. It provides a flexible
alternative to subclassing for extending functionality.

Key components:
1. Component Interface (ReqI): Defines interface for objects that can have responsibilities added
2. Concrete Component (Req): Defines an object to which additional responsibilities can be attached
3. Decorator (LogReqDecorator, MonitorDecorator): Maintains a reference to a Component object and defines an interface that conforms to Component's interface
4. Concrete Decorators: Add responsibilities to the component

Benefits:
- Provides greater flexibility than static inheritance
- Avoids feature-laden classes high up in the hierarchy
- Enhances the object's responsibilities without making code changes
- Supports the Single Responsibility Principle
*/

package main

import (
	"fmt"
	"time"
)

// ReqI defines the interface for request handling components
type ReqI interface {
	Handler(url string) string
}

// Req is the basic request handler that can be decorated
type Req struct {
}

// Handler implements the basic request handling functionality
func (r Req) Handler(url string) string {
	fmt.Println("请求 " + url) // Requesting URL
	time.Sleep(100 * time.Millisecond)
	return ""
}

// LogReqDecorator adds logging functionality to the request handler
type LogReqDecorator struct {
	req ReqI // Reference to the wrapped request handler
}

// Handler adds logging before and after the request handling
func (l LogReqDecorator) Handler(url string) string {
	fmt.Println("日志记录前") // Before logging
	res := l.req.Handler(url)
	fmt.Println("日志记录后") // After logging
	return res
}

// MonitorDecorator adds performance monitoring to the request handler
type MonitorDecorator struct {
	req ReqI // Reference to the wrapped request handler
}

// Handler adds timing measurement to the request handling
func (l MonitorDecorator) Handler(url string) string {
	t1 := time.Now()
	res := l.req.Handler(url)
	sub := time.Since(t1)
	fmt.Println("耗时", sub) // Time elapsed
	return res
}

// Example usage of the Decorator Pattern
func main() {
	// Create the base component
	req := Req{}
	// Add logging decoration
	req1 := LogReqDecorator{req: req}
	// Add monitoring decoration
	req2 := MonitorDecorator{req: req1}
	// Execute the decorated request handler
	req2.Handler("baidu.com")
}
