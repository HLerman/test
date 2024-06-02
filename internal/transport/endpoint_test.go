package transport_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HLerman/test/cmd/api/setup"
	"github.com/HLerman/test/cmd/middleware"
	"github.com/HLerman/test/cmd/provider"
	"github.com/HLerman/test/internal/transport"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func setupRouter() *gin.Engine {
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

	transport.UserHandlers(routeur, setup.SetupUserService(db))
	transport.TransactionHandlers(routeur, setup.SetupTransactionService(db))
	transport.InvoiceHandlers(routeur, setup.SetupInvoiceService(db))

	return routeur
}

func TestEndpoint_GetUsers(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {

		routeur := setupRouter()

		req, _ := http.NewRequest("GET", transport.UsersPath, nil)
		w := httptest.NewRecorder()

		routeur.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		require.Equal(t, `[{"balance":241817,"first_name":"Bob","last_name":"Loco","user_id":1},{"balance":82540,"first_name":"Lynne","last_name":"Gwafranca","user_id":3},{"balance":402758,"first_name":"Art","last_name":"Decco","user_id":4},{"balance":226777,"first_name":"Lynne","last_name":"Gwistic","user_id":5},{"balance":144970,"first_name":"Polly","last_name":"Ester Undawair","user_id":6},{"balance":205387,"first_name":"Oscar","last_name":"Nommanee","user_id":7},{"balance":520060,"first_name":"Laura","last_name":"Biding","user_id":8},{"balance":565074,"first_name":"Laura","last_name":"Norda","user_id":9},{"balance":436180,"first_name":"Des","last_name":"Ignayshun","user_id":10},{"balance":818313,"first_name":"Mike","last_name":"Rowe-Soft","user_id":11},{"balance":189588,"first_name":"Anne","last_name":"Kwayted","user_id":12},{"balance":97005,"first_name":"Wayde","last_name":"Thabalanz","user_id":13},{"balance":276296,"first_name":"Dee","last_name":"Mandingboss","user_id":14},{"balance":932505,"first_name":"Sly","last_name":"Meedentalfloss","user_id":15},{"balance":500691,"first_name":"Stanley","last_name":"Knife","user_id":16},{"balance":478333,"first_name":"Wynn","last_name":"Dozeaplikayshun","user_id":17},{"balance":50298,"first_name":"Kevin","last_name":"Findus","user_id":2}]`, w.Body.String())
	})
}

func TestEndpoint_AddInvoice(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {

		routeur := setupRouter()

		requestBody := gin.H{
			"user_id": 17,
			"amount":  113.45,
			"label":   "Work for April",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", transport.InvoicePath, bytes.NewBuffer(requestBodyBytes))
		w := httptest.NewRecorder()

		routeur.ServeHTTP(w, req)

		require.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("Ko user not found", func(t *testing.T) {

		routeur := setupRouter()

		requestBody := gin.H{
			"user_id": 21,
			"amount":  113.45,
			"label":   "Work for April",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", transport.InvoicePath, bytes.NewBuffer(requestBodyBytes))
		w := httptest.NewRecorder()

		routeur.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
		require.Equal(t, `{"message":"l'identifiant de l'utilisateur est invalide"}`, w.Body.String())
	})

	t.Run("Ko invalid amount", func(t *testing.T) {

		routeur := setupRouter()

		requestBody := gin.H{
			"user_id": 17,
			"amount":  113.455,
			"label":   "Work for April",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", transport.InvoicePath, bytes.NewBuffer(requestBodyBytes))
		w := httptest.NewRecorder()

		routeur.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
		require.Equal(t, `{"message":"le montant invalide"}`, w.Body.String())
	})
}

func TestEndpoint_AddTransaction(t *testing.T) {
	t.Run("Ko invoice not found", func(t *testing.T) {

		routeur := setupRouter()

		requestBody := gin.H{
			"invoice_id": 99,
			"amount":     113.45,
			"reference":  "JMPINV200220117",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", transport.TransactionPath, bytes.NewBuffer(requestBodyBytes))
		w := httptest.NewRecorder()

		routeur.ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
		require.Equal(t, `{"message":"la facture n'existe pas"}`, w.Body.String())
	})
}
