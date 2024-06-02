package models

type Invoice struct {
	Id     int
	Amount int
	Label  string
	UserId int
	Status string
}

const (
	InvoiceStatusPending = "pending"
	InvoiceStatusPaid    = "paid"
)
