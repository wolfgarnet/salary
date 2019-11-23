package salary

// Account is a representation of an account that has an ID and a name
type Account struct {
	// ID is the id of the account
	ID int `json:"id"`

	// Name is the name of the account
	Name string `json:"name"`
}

// TableName returns the table name of account
func (a Account) TableName() string {
	return "account"
}

// ListAccount is a helper struct for representing an account with and accumulated amount
type ListAccount struct {
	// ID is the id of the account
	ID int

	// Name is the name of the account
	Name string

	// Amount is the accumulated amount of all the transactions for the account
	Amount float32
}

// Transactions will return all the transaction for the given account
func (a Account) Transactions() (transactions []Transaction) {
	db.Where(&Transaction{Account:a.ID}).Find(&transactions)
	return
}
