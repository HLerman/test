package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/HLerman/test/internal/business/invoice"
	"github.com/HLerman/test/internal/business/models"
	"github.com/HLerman/test/internal/business/transaction"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	StatusPending = "pending"
	StatusPaid    = "paid"

	GetUsers = `
	SELECT id, first_name, last_name, balance FROM users`

	GetInvoice = `
	SELECT id,
	user_id, status, label, amount FROM invoices WHERE id = $1`

	AddInvoice = `
	INSERT INTO invoices (user_id, status, label, amount)
	VALUES ($1, $2, $3, $4)`

	UpdateUserBalance = `
		UPDATE users SET balance = balance + $1 WHERE id = $2`

	UpdateInvoiceStatus = `
		UPDATE invoices SET status = $1 WHERE id = $2`
)

//go:generate mockgen -source=$GOFILE -package db -destination=db_mock.go

type Db struct {
	db Storage
}

type Storage interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func NewDb(storage Storage) Db {
	return Db{
		db: storage,
	}
}

func (d Db) GetUsers(ctx context.Context) ([]models.User, error) {
	var res []models.User

	rows, err := d.db.QueryContext(ctx, GetUsers)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Balance); err != nil {
			return res, err
		}

		res = append(res, user)
	}

	if err := rows.Err(); err != nil {
		return res, err
	}

	return res, err
}

func (d Db) AddInvoice(ctx context.Context, i models.Invoice) error {
	_, err := d.db.ExecContext(ctx, AddInvoice, i.UserId, StatusPending, i.Label, i.Amount)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation":
				log.Ctx(ctx).Err(err).Msg("invalid user id")
				return invoice.ErrInvalidUserId
			}
		}
		return err
	}

	return nil
}

func (d Db) GetInvoiceByID(ctx context.Context, id int) (*models.Invoice, error) {
	var res models.Invoice

	err := d.db.QueryRowContext(ctx, GetInvoice, id).Scan(
		&res.Id, &res.UserId, &res.Status, &res.Label, &res.Amount,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, transaction.ErrInvoiceNotFound
	case err == nil:
		return &res, nil
	default:
		log.Ctx(ctx).Error().Msg("unable to get information from db")
	}

	return nil, err
}

// AddTransaction implements transaction.Db.
func (d Db) UpdateBalance(ctx context.Context, id int, amount int) error {
	_, err := d.db.ExecContext(ctx, UpdateUserBalance, amount, id)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("unable to update user balance")
		return err
	}
	return nil
}

func (d Db) UpdateInvoice(ctx context.Context, id int) error {
	_, err := d.db.ExecContext(ctx, UpdateInvoiceStatus, StatusPaid, id)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("unable to update invoice status")
		return err
	}
	return nil
}
