package salary

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// System is implements all the top level system functionality
type System struct {
	// db is a reference to the database
	db *gorm.DB
}

// ListAccounts finds all the accounts and accumulated all the transaction amounts.
// The accounts are send through the provided channel as ListAccounts.
func (s *System) ListAccounts(sink chan<- ListAccount) {
	defer close(sink)
	rows, err := s.db.Table("account a").Select("a.id, a.name, sum(t.amount) as amount").Joins("left join 'transaction' t on t.account = a.id").Group("a.id").Rows()
	if err != nil {
		log.Printf("Recieved error while reading data: %v\n", err)
		return
	}
	for rows.Next() {
		var account ListAccount
		rows.Scan(&account.ID, &account.Name, &account.Amount)

		sink <-account
	}
}

// CreateAccount creates an account given a name.
func (s *System) CreateAccount(name string) Account {
	account := Account{
		Name:name,
	}

	s.db.Table("account").Create(&account)
	return account
}

// GetAccount gets an account given an id
func (s *System) GetAccount(id int) (account Account, err error) {
	err = s.db.Where("id = ?", id).First(&account).Error
	return
}

// CreateTransaction creates a transaction given an account, a description and an amount.
func (s *System) CreateTransaction(account *Account, description string, amount float32) (transaction Transaction) {
	transaction = Transaction{
		Account:     account.ID,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}

	s.db.Create(&transaction)
	return
}

// DeleteTransaction deletes a transaction given an id.
func (s *System) DeleteTransaction(id int) {
	s.db.Where("id = ?", id).Delete(Transaction{})
}
