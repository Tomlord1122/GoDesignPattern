/*
Command Pattern

This pattern encapsulates a request as an object, thereby letting you parameterize clients
with different requests, queue or log requests, and support undoable operations.
This implementation demonstrates a task queue that can handle different types of commands
like printing messages, sending emails, and sending text messages.

Key components:
1. Command Interface: Declares interface for executing operation
2. Concrete Commands: Implements Command for specific operations
3. Invoker (TaskQueue): Asks the command to carry out the request
4. Client: Creates and configures concrete Command objects

Benefits:
- Decouples classes that invoke operations from classes that perform operations
- Allows you to create sequences of commands
- Makes it easy to add new commands without changing existing code
- Supports undo/redo operations and command logging
*/

package main

import "fmt"

// Command defines the interface for executing operations
type Command interface {
	Execute()
}

// PrintCommand implements Command for printing messages
type PrintCommand struct {
	Content string
}

// Execute prints the message content
func (p PrintCommand) Execute() {
	fmt.Println("打印消息", p.Content) // Print message
}

// SendEmail implements Command for sending emails
type SendEmail struct {
	To      string // Email recipient
	Content string // Email content
}

// Execute sends an email
func (s SendEmail) Execute() {
	fmt.Println("发送邮件", s.To, s.Content) // Send email
}

// SendTel implements Command for sending text messages
type SendTel struct {
	To      string // Phone number recipient
	Content string // Message content
}

// Execute sends a text message
func (s SendTel) Execute() {
	fmt.Println("发送短信", s.To, s.Content) // Send text message
}

// TaskQueue acts as the invoker that manages and executes commands
type TaskQueue struct {
	Queue []Command // List of commands to execute
}

// NewTaskQueue creates a new task queue
func NewTaskQueue() *TaskQueue {
	return &TaskQueue{}
}

// AddCommand adds a new command to the queue
func (t *TaskQueue) AddCommand(command Command) {
	t.Queue = append(t.Queue, command)
}

// Command executes all commands in the queue
func (t *TaskQueue) Command() {
	for _, command := range t.Queue {
		command.Execute()
	}
}

// Example usage of the Command Pattern
func main() {
	// Create the command invoker
	queue := NewTaskQueue()

	// Add different types of commands to the queue
	queue.AddCommand(&PrintCommand{
		Content: "你好", // Hello
	})
	queue.AddCommand(&SendEmail{
		Content: "你好", // Hello
		To:      "xxx@qq.com",
	})
	queue.AddCommand(&SendTel{
		Content: "你好", // Hello
		To:      "11122223333",
	})

	// Execute all commands in the queue
	queue.Command()
}
