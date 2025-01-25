/*
Observer Pattern

This pattern defines a one-to-many dependency between objects so that when one
object changes state, all its dependents are notified and updated automatically.
This implementation demonstrates a weather station that notifies registered users
of temperature changes.

Key components:
1. Subject Interface: Defines operations for attaching and notifying observers
2. Concrete Subject (WeatherZhan): Maintains state and sends notifications
3. Observer Interface: Defines how observers receive updates
4. Concrete Observer (User): Implements the Observer interface to receive updates

Benefits:
- Supports loose coupling between objects
- Allows sending data to many objects efficiently
- Enables dynamic relationships between objects
- Provides support for broadcast communication
*/

package main

import "fmt"

// Observer defines the interface for objects that should be notified of changes
type Observer interface {
	RevMsg(wd int) // Receive temperature update
}

// User represents a concrete observer that receives weather updates
type User struct {
	Name string // User's name
}

// RevMsg implements the Observer interface for User
func (u User) RevMsg(wd int) {
	fmt.Printf("%s 现在温度%d\n", u.Name, wd) // [Name] current temperature is [wd]
}

// Subject defines the interface for objects that can be observed
type Subject interface {
	SendMsg(wd int)            // Send update to all observers
	Notify()                   // Notify all observers
	RegisterObserver(Observer) // Register new observer
}

// WeatherZhan represents a weather station that monitors temperature
type WeatherZhan struct {
	observerList []Observer // List of registered observers
	wd           int        // Current temperature
}

// SendMsg updates the temperature and notifies all observers
func (w *WeatherZhan) SendMsg(wd int) {
	w.wd = wd
	w.Notify()
}

// Notify sends the current temperature to all registered observers
func (w *WeatherZhan) Notify() {
	for _, observer := range w.observerList {
		observer.RevMsg(w.wd)
	}
}

// RegisterObserver adds a new observer to the notification list
func (w *WeatherZhan) RegisterObserver(observer Observer) {
	w.observerList = append(w.observerList, observer)
}

// Example usage of the Observer Pattern
func main() {
	// Create the subject (weather station)
	zhan := WeatherZhan{}

	// Create observers (users)
	u1 := User{Name: "枫枫"} // User Feng Feng
	u2 := User{Name: "张三"} // User Zhang San

	// Register observers with the subject
	zhan.RegisterObserver(u1)
	zhan.RegisterObserver(u2)

	// Send temperature updates
	zhan.SendMsg(8) // Temperature is 8 degrees
	zhan.SendMsg(0) // Temperature is 0 degrees
}
