/*
Proxy Pattern

This pattern provides a surrogate or placeholder for another object to control access to it.
It's useful for adding a level of indirection for accessing an object, which allows you
to add behavior before or after accessing the real object.

Key components:
1. Subject Interface (implicitly defined by method signatures)
2. Real Subject (Log): The real object that the proxy represents
3. Proxy (ProxyLog): Controls access to the real subject and may be responsible for its creation and lifecycle

Benefits:
- Supports lazy initialization (virtual proxy)
- Adds security control (protection proxy)
- Enables logging and monitoring (logging proxy)
- Facilitates remote resource access (remote proxy)
*/

package main

import "fmt"

// Log represents the real subject that performs the actual logging
type Log struct {
}

// Info performs the actual logging operation
func (Log) Info(content string) {
	fmt.Println("日志落库") // Store log in database
}

// ProxyLog acts as a proxy for the Log object
type ProxyLog struct {
	log *Log // Reference to the real subject
}

// Info proxies the log info call with additional behavior
func (p *ProxyLog) Info(content string) {
	// Lazy initialization of the real subject
	if p.log == nil {
		p.log = &Log{}
	}

	// Pre-processing before the real call
	fmt.Println("Before logging") // Added behavior before the call

	// Actual method call on the real subject
	p.log.Info(content)

	// Post-processing after the real call
	fmt.Println("记录 调用了 info") // Record that Info was called
}

// Example usage of the Proxy Pattern
func main() {
	// Create the proxy without initializing the real subject
	log := ProxyLog{}

	// Use the proxy to call the Info method
	log.Info("zxxx")
}
