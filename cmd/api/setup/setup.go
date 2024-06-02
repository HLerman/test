package setup

import (
	"github.com/HLerman/test/cmd/middleware"
	"github.com/HLerman/test/cmd/provider"
	"github.com/HLerman/test/internal/business/invoice"
	"github.com/HLerman/test/internal/business/transaction"
	"github.com/HLerman/test/internal/business/user"
	"github.com/HLerman/test/internal/transport"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	routeur := gin.New()
	routeur.Use(gin.Recovery())
	routeur.Use(middleware.Logger())

	db := provider.ProvidePostgres(provider.PostgresConf{
		User:     "jump",
		Password: "password",
		Database: "jump",
		Url:      "127.0.0.1",
		Port:     "5432",
	})

	transport.UserHandlers(routeur, SetupUserService(db))
	transport.TransactionHandlers(routeur, SetupTransactionService(db))
	transport.InvoiceHandlers(routeur, SetupInvoiceService(db))

	return routeur
}

func SetupUserService(db user.Db) user.Service {
	return user.NewService(db)
}

func SetupTransactionService(db transaction.Db) transaction.Service {
	return transaction.NewService(db)
}

func SetupInvoiceService(db invoice.Db) invoice.Service {
	return invoice.NewService(db)
}
