package user

import (
	"context"

	"github.com/HLerman/test/internal/business/models"
)

//go:generate mockgen -source=$GOFILE -package user -destination=mock.go

type Db interface {
	GetUsers(ctx context.Context) ([]models.User, error)
}
