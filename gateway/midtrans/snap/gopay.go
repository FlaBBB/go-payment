package snap

import (
	"github.com/midtrans/midtrans-go/snap"

	"github.com/FlaBBB/go-payment/invoice"
)

func NewGopay(inv *invoice.Invoice) (*snap.Request, error) {
	return newBuilder(inv).
		AddPaymentMethods(snap.PaymentTypeGopay).
		Build()
}
