/*
Adapter Pattern

This pattern allows incompatible interfaces to work together by wrapping an object
in an adapter to make it compatible with another class. It's particularly useful
when integrating new systems with existing ones.

Key components:
1. Target Interface (PayI): The interface that clients expect or will use
2. Adaptee (AliPay): The existing class with incompatible interface
3. Adapter (AliPayAdapter): Adapts the Adaptee to the Target interface
4. Client: Collaborates with objects conforming to the Target interface

Benefits:
- Allows incompatible interfaces to work together
- Increases reusability of existing code
- Provides flexibility to add new adapters
- Separates interface conversion from business logic
*/

package main

import "fmt"

// AliPay represents the Adaptee - the existing payment system with incompatible interface
type AliPay struct {
}

// GetPayPage is the existing method with incompatible interface
func (a AliPay) GetPayPage(price int64) string {
	return "支付宝支付链接" // Alipay payment link
}

// AliPayAdapter adapts AliPay to match the PayI interface
type AliPayAdapter struct {
	aliPay *AliPay
}

// PayPage adapts the GetPayPage method to match the PayI interface
func (w AliPayAdapter) PayPage(price int64) string {
	return w.aliPay.GetPayPage(price)
}

// WeiXinPay represents a payment system that already matches the target interface
type WeiXinPay struct {
}

// PayPage implements the PayI interface directly
func (w WeiXinPay) PayPage(price int64) string {
	return "微信支付链接" // WeChat payment link
}

// PayI defines the target interface that clients expect
type PayI interface {
	PayPage(price int64) string
}

// PayPage is a client function that works with the PayI interface
func PayPage(pi PayI, price int64) string {
	return pi.PayPage(price)
}

// Example usage of the Adapter Pattern
func main() {
	// Using WeiXinPay which already implements PayI
	fmt.Println(PayPage(WeiXinPay{}, 1))

	// Using AliPay through the adapter to match PayI
	fmt.Println(PayPage(AliPayAdapter{aliPay: &AliPay{}}, 1))
}
