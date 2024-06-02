package invoice

import (
	"context"

	"github.com/HLerman/test/internal/business/models"
)

//go:generate mockgen -source=$GOFILE -package invoice -destination=mock.go

type Db interface {
	AddInvoice(ctx context.Context, invoice models.Invoice) error
}
