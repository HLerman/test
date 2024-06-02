package db_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HLerman/test/internal/business/invoice"
	"github.com/HLerman/test/internal/business/models"
	"github.com/HLerman/test/internal/implementation/db"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer sqlDb.Close()

	d := db.NewDb(sqlDb)

	t.Run("ok", func(t *testing.T) {
		expected := []models.User{
			{
				Balance:   100,
				FirstName: "John",
				LastName:  "Doeuf",
				UserId:    123,
			},
			{
				Balance:   200,
				FirstName: "Alain",
				LastName:  "Proviste",
				UserId:    124,
			},
		}

		rows := sqlmock.NewRows([]string{
			"id",
			"first_name",
			"last_name",
			"balance",
		}).AddRow(123, "John", "Doeuf", 100).
			AddRow(124, "Alain", "Proviste", 200)

		mock.ExpectQuery(db.GetUsers).WillReturnRows(rows)

		users, err := d.GetUsers(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expected, users)
	})

	t.Run("db_error", func(t *testing.T) {
		mock.ExpectQuery(db.GetUsers).WillReturnError(sql.ErrConnDone)

		_, err := d.GetUsers(context.Background())

		assert.Equal(t, err, sql.ErrConnDone)
	})
}

func TestAddInvoice(t *testing.T) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer sqlDb.Close()

	d := db.NewDb(sqlDb)

	t.Run("ko foreign key", func(t *testing.T) {
		pqErr := &pq.Error{
			Code:    "23503",
			Message: "foreign_key_violation",
		}
		mock.ExpectExec("INSERT INTO invoices").WithArgs(20, "pending", "label", 100).WillReturnError(pqErr)

		err := d.AddInvoice(context.Background(), models.Invoice{
			Amount: 100,
			Label:  "label",
			UserId: 20,
		})

		assert.ErrorIs(t, err, invoice.ErrInvalidUserId)
	})

	t.Run("ok", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO invoices").WithArgs(20, "pending", "label", 100).WillReturnResult(sqlmock.NewResult(20, 1))
		err := d.AddInvoice(context.Background(), models.Invoice{
			Amount: 100,
			Label:  "label",
			UserId: 20,
		})

		assert.NoError(t, err)
	})
}

func TestUpdateBalance(t *testing.T) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer sqlDb.Close()

	d := db.NewDb(sqlDb)

	t.Run("ok", func(t *testing.T) {
		mock.ExpectExec("UPDATE users").WithArgs(100, 20).WillReturnResult(sqlmock.NewResult(20, 1))
		err := d.UpdateBalance(context.Background(), 20, 100)

		assert.NoError(t, err)
	})
}

func TestUpdateInvoice(t *testing.T) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer sqlDb.Close()

	d := db.NewDb(sqlDb)

	t.Run("ok", func(t *testing.T) {
		mock.ExpectExec("UPDATE invoices").WithArgs("paid", 20).WillReturnResult(sqlmock.NewResult(20, 1))
		err := d.UpdateInvoice(context.Background(), 20)

		assert.NoError(t, err)
	})
}
