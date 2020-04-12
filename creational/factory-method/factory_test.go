package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Errorf("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(100.50)
	if !strings.Contains(msg, "paid using cash") {
		t.Errorf("The cash payment method message wasn't correct")
	}
	t.Log("Log: ", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Errorf("A payment method of type 'Debit Card' must exist")
	}

	msg := payment.Pay(100.50)
	if !strings.Contains(msg, "paid using debit card") {
		t.Errorf("The cash payment method message wasn't correct")
	}
	t.Log("Log: ", msg)
}
