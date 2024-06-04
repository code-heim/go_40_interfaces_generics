package main

import "fmt"

// PaymentProcessor defines the behavior for processing payments
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// CreditCard struct
type CreditCard struct {
	CardNumber string
}

// ProcessPayment method for CreditCard
func (cc CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf(
		"Processed payment of $%.2f using Credit Card %s",
		amount,
		cc.CardNumber)
}

// PayPal struct
type PayPal struct {
	Email string
}

// ProcessPayment method for PayPal
func (pp PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf(
		"Processed payment of $%.2f using PayPal account %s",
		amount,
		pp.Email)
}

// BankTransfer struct
type BankTransfer struct {
	AccountNumber string
}

// ProcessPayment method for BankTransfer
func (bt BankTransfer) ProcessPayment(amount float64) string {
	return fmt.Sprintf(
		"Processed payment of $%.2f using Bank Transfer %s",
		amount,
		bt.AccountNumber)
}

// process function uses the PaymentProcessor interface
func process(payment PaymentProcessor, amount float64) {
	fmt.Println(payment.ProcessPayment(amount))
}

func main() {
	cc := CreditCard{CardNumber: "1234-5678-8765-4444"}
	pp := PayPal{Email: "user1@codeheim.io"}
	bt := BankTransfer{AccountNumber: "987654321"}

	process(cc, 100.95)
	process(pp, 170.75)
	process(bt, 250.50)
}
