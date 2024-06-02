package transport

import (
	"context"
	"net/http"

	api "github.com/HLerman/test"
	"github.com/HLerman/test/cmd/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=$GOFILE -package transport -destination=mock.go

type InvoiceService interface {
	AddInvoice(ctx context.Context, invoice api.PostInvoiceJSONRequestBody) error
}

type TransactionService interface {
	AddTransaction(ctx context.Context, transaction api.PostTransactionJSONRequestBody) error
}

type UserService interface {
	GetUsers(ctx context.Context) (api.UserArray, error)
}

func TransactionHandlers(r *gin.Engine, s TransactionService) {
	r.POST(TransactionPath, addTransaction(s))
}

func InvoiceHandlers(r *gin.Engine, s InvoiceService) {
	r.POST(InvoicePath, addInvoice(s))
}

func UserHandlers(r *gin.Engine, s UserService) {
	r.GET(UsersPath, getUsers(s))
}

func getUsers(s UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(middleware.GetLogger(c.Request.Context()).WithContext(c.Request.Context()))
		resp, err := s.GetUsers(c.Request.Context())
		if err != nil {
			ErrorReponse(c, err)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func addInvoice(s InvoiceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(middleware.GetLogger(c.Request.Context()).WithContext(c.Request.Context()))

		var req api.PostInvoiceJSONRequestBody
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Err(err).Msg("failed binding")
			BindingErrorResponse(c, err)
			return
		}

		err := s.AddInvoice(c.Request.Context(), req)
		if err != nil {
			ErrorReponse(c, err)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func addTransaction(s TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(middleware.GetLogger(c.Request.Context()).WithContext(c.Request.Context()))

		var req api.PostTransactionJSONRequestBody
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Err(err).Msg("failed binding")
			BindingErrorResponse(c, err)
			return
		}

		err := s.AddTransaction(c.Request.Context(), req)
		if err != nil {
			ErrorReponse(c, err)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
