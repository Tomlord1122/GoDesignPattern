/*
Factory Method Pattern

This pattern defines an interface for creating objects but lets subclasses decide which class to instantiate.
It's particularly useful when a class can't anticipate the type of objects it needs to create.

Key components:
1. Product Interface (Pay): Defines the interface for objects the factory method creates
2. Concrete Products (AliPay, WeiXinPay, YinLianPay): Different implementations of the product interface
3. Creator Interface (PayFactory): Declares the factory method
4. Concrete Creators (AliPayFactory, WeiXinPayFactory, YinLianPayFactory): Implement the factory method

Benefits:
- Follows Single Responsibility Principle by moving product creation code into one place
- Follows Open/Closed Principle as you can introduce new types of products without breaking existing code
- Provides hooks for subclasses to extend the creation process
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

// YinLianPay implements the Pay interface for UnionPay payment method
type YinLianPay struct {
}

// PayPage implements the payment page display for UnionPay
func (YinLianPay) PayPage(price int64) string {
	return "银联支付"
}

// PayFactory defines the interface for creating payment objects
type PayFactory interface {
	CreatePay() Pay
}

// AliPayFactory creates Alipay payment instances
type AliPayFactory struct {
}

// CreatePay implements the factory method for creating Alipay instances
func (AliPayFactory) CreatePay() Pay {
	// Contains Alipay-specific creation logic
	return AliPay{}
}

// WeiXinPayFactory creates WeChat payment instances
type WeiXinPayFactory struct {
}

// CreatePay implements the factory method for creating WeChat Pay instances
func (WeiXinPayFactory) CreatePay() Pay {
	return WeiXinPay{}
}

// YinLianPayFactory creates UnionPay payment instances
type YinLianPayFactory struct {
}

// CreatePay implements the factory method for creating UnionPay instances
func (YinLianPayFactory) CreatePay() Pay {
	return YinLianPay{}
}

// Example usage of the Factory Method Pattern
func main() {
	// Create and use Alipay payment
	aliPayFactory := AliPayFactory{}
	aliPay := aliPayFactory.CreatePay()
	fmt.Println(aliPay.PayPage(15))

	// Create and use WeChat payment
	weiXinPayFactory := WeiXinPayFactory{}
	weiXinPay := weiXinPayFactory.CreatePay()
	fmt.Println(weiXinPay.PayPage(15))

	// Create and use UnionPay payment
	yinLianPayFactory := YinLianPayFactory{}
	yinLianPay := yinLianPayFactory.CreatePay()
	fmt.Println(yinLianPay.PayPage(15))
}
