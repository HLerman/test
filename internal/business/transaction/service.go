package transaction

import (
	"context"

	api "github.com/HLerman/test"
	"github.com/HLerman/test/internal/business/models"
	"github.com/HLerman/test/internal/libs/amount"
)

type Service struct {
	db Db
}

func NewService(db Db) Service {
	return Service{
		db: db,
	}
}

func (s Service) AddTransaction(ctx context.Context, transaction api.PostTransactionJSONRequestBody) error {
	cents, err := amount.ConvertToCents(transaction.Amount)
	if err != nil {
		return err
	}

	invoice, err := s.db.GetInvoiceByID(ctx, transaction.InvoiceId)
	if err != nil {
		return err
	}

	if invoice.Status == models.InvoiceStatusPaid {
		return ErrInvoiceAlreadyPayed
	}

	if invoice.Amount != cents {
		return ErrTransactionAmountDoesntFeatInvoiceAmount
	}

	err = s.db.UpdateBalance(ctx, invoice.Amount, invoice.UserId)
	if err != nil {
		return err
	}

	err = s.db.UpdateInvoice(ctx, invoice.Id)
	if err != nil {
		return err
	}

	return nil
}
