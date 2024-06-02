package transport

import (
	"net/http"

	api "github.com/HLerman/test"
	"github.com/HLerman/test/internal/business/invoice"
	"github.com/HLerman/test/internal/business/transaction"
	"github.com/HLerman/test/internal/libs/amount"
	"github.com/gin-gonic/gin"
)

type Error struct {
	StatusCode int
	Err        api.Error
}

var mappedError = map[error]Error{
	transaction.ErrTransactionAmountDoesntFeatInvoiceAmount: {
		StatusCode: http.StatusBadRequest,
		Err: api.Error{
			Message: transaction.ErrTransactionAmountDoesntFeatInvoiceAmount.Error(),
		},
	},
	transaction.ErrInvoiceAlreadyPayed: {
		StatusCode: http.StatusUnprocessableEntity,
		Err: api.Error{
			Message: transaction.ErrInvoiceAlreadyPayed.Error(),
		},
	},
	transaction.ErrInvoiceNotFound: {
		StatusCode: http.StatusNotFound,
		Err: api.Error{
			Message: transaction.ErrInvoiceNotFound.Error(),
		},
	},
	amount.ErrInvalidAmount: {
		StatusCode: http.StatusBadRequest,
		Err: api.Error{
			Message: amount.ErrInvalidAmount.Error(),
		},
	},
	invoice.ErrInvalidUserId: {
		StatusCode: http.StatusBadRequest,
		Err: api.Error{
			Message: invoice.ErrInvalidUserId.Error(),
		},
	},
}

func ErrorReponse(c *gin.Context, e error) {
	if v, ok := mappedError[e]; ok {
		c.JSON(v.StatusCode, v.Err)
		return
	}

	c.JSON(http.StatusInternalServerError, api.Error{Message: "internal server error"})
}

func BindingErrorResponse(c *gin.Context, e error) {
	c.JSON(http.StatusBadRequest, api.Error{Message: e.Error()})
}
