package salary

import "time"

type Transaction struct {
	ID int `json:"id"`
	//Account *Account `json:"account"`
	Account int `json:"account"`
	Description string `json:"description"`
	Amount float32 `json:"amount"`
	Date time.Time `json:"date"`
}

func (t Transaction) TableName() string {
	return "transaction"
}
