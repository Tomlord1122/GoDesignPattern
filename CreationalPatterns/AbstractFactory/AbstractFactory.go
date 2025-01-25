/*
Abstract Factory Pattern

This pattern provides an interface for creating families of related or dependent objects
without specifying their concrete classes. It's particularly useful when a system needs to be
independent from how its products are created, composed, and represented.

Key components:
1. Abstract Products (Pay, Refund): Interfaces for families of related products
2. Concrete Products (AliPay/WeiXinPay, AliRefund/WeiXinRefund): Implementations of the product interfaces
3. Abstract Factory (PayFactory): Interface declaring creation methods for all abstract products
4. Concrete Factories (AliPayFactory, WeiXinPayFactory): Implement creation methods for concrete products

Benefits:
- Ensures compatibility between related products
- Isolates concrete classes from client code
- Makes exchanging product families easy
- Promotes consistency among products
*/

package main

import "fmt"

// Pay defines the interface for payment processing
type Pay interface {
	PayPage(price int64) string
}

// AliPay implements the Pay interface for Alipay
type AliPay struct {
}

// PayPage implements the payment page display for Alipay
func (AliPay) PayPage(price int64) string {
	return "支付宝支付"
}

// WeiXinPay implements the Pay interface for WeChat Pay
type WeiXinPay struct {
}

// PayPage implements the payment page display for WeChat Pay
func (WeiXinPay) PayPage(price int64) string {
	return "微信支付"
}

// Refund defines the interface for refund processing
type Refund interface {
	Refund(no string) error
}

// AliRefund implements the Refund interface for Alipay refunds
type AliRefund struct {
}

// Refund processes refunds for Alipay transactions
func (AliRefund) Refund(no string) error {
	fmt.Println("支付宝退款啦")
	return nil
}

// WeiXinRefund implements the Refund interface for WeChat Pay refunds
type WeiXinRefund struct {
}

// Refund processes refunds for WeChat Pay transactions
func (WeiXinRefund) Refund(no string) error {
	fmt.Println("微信退款啦")
	return nil
}

// PayFactory defines the abstract factory interface for creating payment-related objects
type PayFactory interface {
	CreatePay() Pay
	CreateRefund() Refund
}

// AliPayFactory creates Alipay-related objects
type AliPayFactory struct {
}

// CreatePay creates an Alipay payment processor
func (AliPayFactory) CreatePay() Pay {
	// Contains Alipay-specific creation logic
	return AliPay{}
}

// CreateRefund creates an Alipay refund processor
func (AliPayFactory) CreateRefund() Refund {
	return AliRefund{}
}

// WeiXinPayFactory creates WeChat Pay-related objects
type WeiXinPayFactory struct {
}

// CreatePay creates a WeChat Pay payment processor
func (WeiXinPayFactory) CreatePay() Pay {
	return WeiXinPay{}
}

// CreateRefund creates a WeChat Pay refund processor
func (WeiXinPayFactory) CreateRefund() Refund {
	return WeiXinRefund{}
}

// Example usage of the Abstract Factory Pattern
func main() {
	// Create and use Alipay family of products
	aliPayFactory := AliPayFactory{}
	aliPay := aliPayFactory.CreatePay()
	fmt.Println(aliPay.PayPage(15))
	aliPayFactory.CreateRefund().Refund("")

	// Create and use WeChat Pay family of products
	weiXinPayFactory := WeiXinPayFactory{}
	weiXinPay := weiXinPayFactory.CreatePay()
	fmt.Println(weiXinPay.PayPage(15))
	weiXinPayFactory.CreateRefund().Refund("")
}
