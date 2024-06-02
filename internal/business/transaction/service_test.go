package transaction_test

import (
	"context"
	"testing"

	api "github.com/HLerman/test"
	"github.com/HLerman/test/internal/business/models"
	"github.com/HLerman/test/internal/business/transaction"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestService_AddTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := transaction.NewMockDb(ctrl)
	svc := transaction.NewService(db)

	t.Run("ok", func(t *testing.T) {
		db.EXPECT().GetInvoiceByID(gomock.Any(), gomock.Any()).Return(&models.Invoice{
			Id:     17,
			Amount: 100,
			Label:  "test",
			UserId: 15,
			Status: "pending",
		}, nil)

		db.EXPECT().UpdateBalance(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		db.EXPECT().UpdateInvoice(gomock.Any(), gomock.Any()).Return(nil)

		err := svc.AddTransaction(context.Background(), api.PostTransactionJSONRequestBody{
			Amount:    1,
			InvoiceId: 17,
			Reference: "test reference",
		})

		require.NoError(t, err)
	})

	t.Run("ko invoice not found", func(t *testing.T) {
		db.EXPECT().GetInvoiceByID(gomock.Any(), gomock.Any()).Return(nil, transaction.ErrInvoiceNotFound)

		err := svc.AddTransaction(context.Background(), api.PostTransactionJSONRequestBody{
			Amount:    1,
			InvoiceId: 17,
			Reference: "test reference",
		})

		require.ErrorIs(t, err, transaction.ErrInvoiceNotFound)
	})

	t.Run("ko invoice already payed", func(t *testing.T) {
		db.EXPECT().GetInvoiceByID(gomock.Any(), gomock.Any()).Return(&models.Invoice{
			Id:     17,
			Amount: 100,
			Label:  "test",
			UserId: 15,
			Status: "paid",
		}, nil)

		err := svc.AddTransaction(context.Background(), api.PostTransactionJSONRequestBody{
			Amount:    1,
			InvoiceId: 17,
			Reference: "test reference",
		})

		require.ErrorIs(t, err, transaction.ErrInvoiceAlreadyPayed)
	})

	t.Run("ko transaction amount doesnt feat invoice amount", func(t *testing.T) {
		db.EXPECT().GetInvoiceByID(gomock.Any(), gomock.Any()).Return(&models.Invoice{
			Id:     17,
			Amount: 100,
			Label:  "test",
			UserId: 15,
			Status: "pending",
		}, nil)

		err := svc.AddTransaction(context.Background(), api.PostTransactionJSONRequestBody{
			Amount:    10,
			InvoiceId: 17,
			Reference: "test reference",
		})

		require.ErrorIs(t, err, transaction.ErrTransactionAmountDoesntFeatInvoiceAmount)
	})
}
