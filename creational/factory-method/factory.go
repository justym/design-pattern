package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	//Cash is PaymentMethod
	Cash = iota
	//DebitCard is PaymentMethod
	DebitCard
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &CashPM{}, nil
	case DebitCard:
		return &DebitCardPM{}, nil
	}

	return nil, errors.New("Unkonw method")
}

type CashPM struct {
	PM PaymentMethod
}

type DebitCardPM struct {
	PM PaymentMethod
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("paid using cash: %g", amount)
}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("paid using debit card: %g", amount)
}
