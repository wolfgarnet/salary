package salary

import "time"

// Transaction represents a transaction for a given account
type Transaction struct {
	// ID is the id of the transaction
	ID int `json:"id"`

	// Account is of the associated account id of the transaction
	Account int `json:"account"`

	// Description is the description of the transaction
	Description string `json:"description"`

	// Amount is the amount of the transaction
	Amount float32 `json:"amount"`

	// Date is the date when the transaction was registered by the system
	Date time.Time `json:"date"`
}

// TableName returns the table name of transaction
func (t Transaction) TableName() string {
	return "transaction"
}
