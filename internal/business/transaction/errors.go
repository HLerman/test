package transaction

import "errors"

var ErrInvoiceAlreadyPayed = errors.New("la facture est déjà payée")
var ErrTransactionAmountDoesntFeatInvoiceAmount = errors.New("le montant de la transaction ne correspond pas au montant de la facture")
var ErrInvoiceNotFound = errors.New("la facture n'existe pas")
