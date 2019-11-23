package salary

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

var (
	db *gorm.DB
	err error
)

func Initialize() *System {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	//defer db.Close()

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


	system := &System{
		db:db,
	}
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

	return system
}

func deleteDatabase() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	log.Printf("Deleting database from %v", dir)
	return os.Remove("./gorm.db")

}
