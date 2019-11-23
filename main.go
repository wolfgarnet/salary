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

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	system := &System{
		db:db,
	}
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
