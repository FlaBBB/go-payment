package xeninvoice

import (
	xinvoice "github.com/xendit/xendit-go/invoice"

	"github.com/FlaBBB/go-payment/invoice"
)

func NewCreditCard(inv *invoice.Invoice) (*xinvoice.CreateParams, error) {
	return newBuilder(inv).
		AddPaymentMethod("CREDIT_CARD").
		Build()
}
