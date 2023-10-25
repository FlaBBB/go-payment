package snap_test

import (
	"testing"

	midsnap "github.com/midtrans/midtrans-go/snap"
	"github.com/stretchr/testify/assert"

	"github.com/FlaBBB/go-payment/gateway/midtrans/snap"
	"github.com/FlaBBB/go-payment/invoice"
)

func TestNewBNIVA(t *testing.T) {
	type args struct {
		inv *invoice.Invoice
	}
	tests := []struct {
		name    string
		args    args
		want    *midsnap.Request
		wantErr error
	}{
		{
			name: "standard bni va request",
			args: args{inv: dummyInv},
			want: &midsnap.Request{
				EnabledPayments: []midsnap.SnapPaymentType{
					midsnap.PaymentTypeBNIVA,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := snap.NewBNIVA(tt.args.inv)
			assert.Equal(t, tt.wantErr, err)
			assert.EqualValues(t, tt.want.BniVa, got.BniVa)
			assert.Contains(t, got.EnabledPayments, midsnap.PaymentTypeBNIVA)
		})
	}
}
