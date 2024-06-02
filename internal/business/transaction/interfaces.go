package transaction

import (
	"context"

	"github.com/HLerman/test/internal/business/models"
)

//go:generate mockgen -source=$GOFILE -package transaction -destination=mock.go

type Db interface {
	UpdateInvoice(ctx context.Context, id int) error
	UpdateBalance(ctx context.Context, id int, amount int) error
	GetInvoiceByID(ctx context.Context, id int) (*models.Invoice, error)
}
