/*
Strategy Pattern

This pattern defines a family of algorithms, encapsulates each one, and makes them
interchangeable. It lets the algorithm vary independently from clients that use it.
This implementation demonstrates different payment methods that can be used
interchangeably.

Key components:
1. Strategy Interface (Pay): Declares interface common to all algorithms
2. Concrete Strategies (AliPay, WeiPay): Implements different algorithms
3. Context (PayStrategy): Maintains reference to strategy and delegates it
4. Client: Creates specific strategies and passes them to context

Benefits:
- Provides alternative to subclassing for algorithmic variation
- Eliminates complex conditional statements
- Allows runtime switching between algorithms
- Promotes reuse of algorithmic code
*/

package main

import "fmt"

// Pay defines the interface for different payment strategies
type Pay interface {
	Pay(money int64) // Execute payment with specified amount
}

// AliPay implements the Alipay payment strategy
type AliPay struct {
}

// Pay processes payment using Alipay
func (AliPay) Pay(money int64) {
	fmt.Println("使用支付宝支付", money) // Using Alipay to pay [amount]
}

// WeiPay implements the WeChat Pay payment strategy
type WeiPay struct {
}

// Pay processes payment using WeChat Pay
func (WeiPay) Pay(money int64) {
	fmt.Println("使用微信支付", money) // Using WeChat Pay to pay [amount]
}

// PayStrategy acts as the context for payment strategies
type PayStrategy struct {
	pay Pay // Current payment strategy
}

// SetPay changes the payment strategy
func (p *PayStrategy) SetPay(pay Pay) {
	p.pay = pay
}

// Pay delegates the payment to the current strategy
func (p *PayStrategy) Pay(money int64) {
	p.pay.Pay(money)
}

// Example usage of the Strategy Pattern
func main() {
	// Create payment strategies
	aliPay := &AliPay{}
	weiPay := &WeiPay{}

	// Create context and use different strategies
	payStrategy := &PayStrategy{}

	// Use Alipay strategy
	payStrategy.SetPay(aliPay)
	payStrategy.Pay(12)

	// Switch to WeChat Pay strategy
	payStrategy.SetPay(weiPay)
	payStrategy.Pay(12)
}
