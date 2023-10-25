package xeninvoice

import (
	xinvoice "github.com/xendit/xendit-go/invoice"

	"github.com/FlaBBB/go-payment/invoice"
)

func NewBRIVA(inv *invoice.Invoice) (*xinvoice.CreateParams, error) {
	return newBuilder(inv).
		AddPaymentMethod("BRI").
		Build()
}
