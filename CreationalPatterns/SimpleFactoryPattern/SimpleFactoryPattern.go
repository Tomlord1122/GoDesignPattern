/*
Simple Factory Pattern

This pattern provides a way to create objects without exposing the creation logic to the client.
It encapsulates the object creation process in a single factory method.

Key components:
1. Product Interface (Pay): Defines the interface for objects the factory will create
2. Concrete Products (AliPay, WeiXinPay): Implementations of the product interface
3. Factory (NewPayPage): Creates and returns product instances based on input parameters

Benefits:
- Encapsulates object creation logic
- Provides a simple way to create different product variants
- Makes the code more maintainable by centralizing object creation
*/

package main

import "fmt"

// Pay defines the interface that all payment methods must implement
type Pay interface {
	PayPage(price int64) string
}

// AliPay implements the Pay interface for Alipay payment method
type AliPay struct {
}

// PayPage implements the payment page display for Alipay
func (AliPay) PayPage(price int64) string {
	return "支付宝支付"
}

// WeiXinPay implements the Pay interface for WeChat payment method
type WeiXinPay struct {
}

// PayPage implements the payment page display for WeChat Pay
func (WeiXinPay) PayPage(price int64) string {
	return "微信支付"
}

// PayType represents different types of payment methods
type PayType int8

const (
	AliPayType    PayType = 1 // Alipay payment type
	WeiXinPayType PayType = 2 // WeChat payment type
)

// NewPayPage is the factory method that creates payment instances
// based on the specified payment type
func NewPayPage(payType PayType) Pay {
	switch payType {
	case AliPayType:
		return AliPay{}
	case WeiXinPayType:
		return WeiXinPay{}
	}
	return nil
}

// Example usage of the Simple Factory Pattern
func main() {
	// Create and use Alipay payment
	pay := NewPayPage(AliPayType)
	fmt.Println(pay.PayPage(12))

	// Create and use WeChat payment
	pay = NewPayPage(WeiXinPayType)
	fmt.Println(pay.PayPage(12))
}
