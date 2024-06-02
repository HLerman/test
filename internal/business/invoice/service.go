package invoice

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

func (s Service) AddInvoice(ctx context.Context, invoice api.PostInvoiceJSONRequestBody) error {
	cents, err := amount.ConvertToCents(invoice.Amount)
	if err != nil {
		return err
	}

	return s.db.AddInvoice(ctx, models.Invoice{
		Label:  invoice.Label,
		UserId: invoice.UserId,
		Amount: cents,
	})
}
