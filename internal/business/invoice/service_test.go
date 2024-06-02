package invoice_test

import (
	"context"
	"testing"

	api "github.com/HLerman/test"
	"github.com/HLerman/test/internal/business/invoice"
	"github.com/HLerman/test/internal/libs/amount"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestService_AddInvoice(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := invoice.NewMockDb(ctrl)
	svc := invoice.NewService(db)

	t.Run("ok", func(t *testing.T) {
		db.EXPECT().AddInvoice(gomock.Any(), gomock.Any()).Return(nil)

		err := svc.AddInvoice(context.Background(), api.PostInvoiceJSONRequestBody{
			Amount: 100.95,
			Label:  "label",
			UserId: 17,
		})

		require.NoError(t, err)
	})

	t.Run("ko invalid amount", func(t *testing.T) {
		err := svc.AddInvoice(context.Background(), api.PostInvoiceJSONRequestBody{
			Amount: 100.955,
			Label:  "label",
			UserId: 17,
		})

		require.ErrorIs(t, err, amount.ErrInvalidAmount)
	})
}
