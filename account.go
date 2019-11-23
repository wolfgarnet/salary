package salary

type Account struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func (a Account) TableName() string {
	return "account"
}

type ListAccount struct {
	ID int
	Name string
	Amount float32
}

func (a Account) Transactions() (transactions []Transaction) {
	db.Where(&Transaction{Account:a.ID}).Find(&transactions)

	return
}
