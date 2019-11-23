package salary

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"time"
)

type System struct {
	db *gorm.DB
}

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

		//fmt.Printf("I HAV22222E %+v\n", account)
		sink <-account
	}
}

func (s *System) CreateAccount(name string) Account {
	account := Account{
		Name:name,
	}

	s.db.Table("account").Create(&account)
	return account
}

func (s *System) GetAccount(id int) (account Account, err error) {
	err = s.db.Where("id = ?", id).First(&account).Error
	return
}

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

func (s *System) CreateRandomTransaction(account *Account) (transaction Transaction) {
	transaction = Transaction{
		Account:     account.ID,
		Description: fmt.Sprintf("Description for %v", account.Name),
		Amount:      rand.Float32()*1000,
		Date:        time.Now(),
	}

	s.db.Create(&transaction)
	return
}

func (s *System) DeleteTransaction(id int) {
	s.db.Where("id = ?", id).Delete(Transaction{})
}
