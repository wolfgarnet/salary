package salary

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
	err error
)

func Initialize() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	/*acc1 := Account{
		Name: "wolle",
	}
	db.Create(&acc1)

	var account Account
	//db.First(&account)
	//fmt.Printf("ACCOUNT: %+v\n", account)
	err := db.Where("id = ?", 1).First(&account).Error
	if err != nil {
		fmt.Printf("DAMN! %v\n", err)
	} else {
		fmt.Printf("ACCOUNT: %v\n", account)
	}*/



	tables := []string{}
	db.Select(&tables, "SHOW TABLES")
	fmt.Println(tables)

	system := System{}
	/*account := system.CreateAccount("knolle")
	fmt.Printf("ACCOUNT: %v\n", account)*/

	/*acc1, err := system.GetAccount(1)
	fmt.Printf("ACCOUNT: %v --- %v\n", acc1, err)*/

	/*transaction := system.CreateRandomTransaction(&account)
	fmt.Printf("TRANSACTION: %+v\n", transaction)*/

	/*transactions := acc1.Transactions()
	fmt.Printf("Transactions: %v\n", transactions)*/

	source := make(chan ListAccount, 128)
	go system.ListAccounts(source)
	for account := range source {
		fmt.Printf("Account: %+v\n", account)
	}
	//fmt.Printf("I AHVE: %v\n", accounts)
}

//db, _ = gorm.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/samples?charset=utf8&parseTime=True&loc=Local")
