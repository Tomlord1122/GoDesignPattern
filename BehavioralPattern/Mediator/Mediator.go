/*
Mediator Pattern

This pattern defines an object that encapsulates how a set of objects interact.
It promotes loose coupling by keeping objects from referring to each other explicitly,
and lets you vary their interaction independently. This implementation demonstrates
a chat room where users can communicate through a mediator.

Key components:
1. Mediator Interface: Defines interface for communicating with Colleague objects
2. Concrete Mediator (ChatRoom): Implements cooperative behavior by coordinating Colleague objects
3. Colleague Interface (Obj): Defines interface for objects communicating through mediator
4. Concrete Colleague (User): Implements objects that communicate through the mediator

Benefits:
- Reduces coupling between components
- Centralizes control of object interactions
- Simplifies object protocols
- Makes object interaction more reusable
*/

package main

import "fmt"

// Obj defines the interface for objects that can communicate through the mediator
type Obj interface {
	SendMsg(string) // Send a message
	RevMsg(string)  // Receive a message
}

// Mediator defines the interface for the communication mediator
type Mediator interface {
	SendMsg(msg string, user Obj) // Distribute message to other users
}

// User represents a chat participant
type User struct {
	Name     string   // User's name
	mediator Mediator // Reference to the mediator
}

// SendMsg sends a message through the mediator
func (u User) SendMsg(msg string) {
	fmt.Printf("用户 %s 发了消息 %s\n", u.Name, msg) // User [name] sent message [msg]
	u.mediator.SendMsg(msg, u)
}

// RevMsg receives a message from the mediator
func (u User) RevMsg(msg string) {
	fmt.Printf("用户 %s 接收到消息 %s\n", u.Name, msg) // User [name] received message [msg]
}

// ChatRoom implements the mediator for user communication
type ChatRoom struct {
	users []User // List of users in the chat room
}

// Register adds a new user to the chat room
func (c *ChatRoom) Register(user User) {
	c.users = append(c.users, user)
}

// SendMsg distributes a message to all users except the sender
func (c *ChatRoom) SendMsg(msg string, user Obj) {
	for _, u := range c.users {
		if u == user {
			continue // Skip the sender
		}
		u.RevMsg(msg)
	}
}

// Example usage of the Mediator Pattern
func main() {
	// Create the mediator (chat room)
	room := ChatRoom{}

	// Create users with references to the mediator
	u1 := User{Name: "枫枫", mediator: &room} // User Feng Feng
	u2 := User{Name: "张三", mediator: &room} // User Zhang San
	u3 := User{Name: "李四", mediator: &room} // User Li Si

	// Register users in the chat room
	room.Register(u1)
	room.Register(u2)
	room.Register(u3)

	// Demonstrate communication through the mediator
	u1.SendMsg("你好啊") // Hello
	u2.SendMsg("吃了吗") // Have you eaten?
	u3.SendMsg("我吃了") // I have eaten
}
