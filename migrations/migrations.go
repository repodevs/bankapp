package migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/repodevs/bankapp/helpers"
)

// User struct for user
type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}

// Account struct for account
type Account struct {
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint
}


func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "postgres://postgres:123456@localhost:5432/go_bankapp?sslmode=disable")
	helpers.HandleErr(err)
	return db
}


// createAccounts used for create dummy account
func createAccounts() {
	db := connectDB()

	defer db.Close()

	users := [2]User{
		{Username: "Anton", Email: "Anton@bank.app"},
		{Username: "Bayu", Email: "Bayu@bank.app"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Debug().Create(&user)

		account := Account{Type: "daily account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Debug().Create(&account)
	}
}

// Migrate create automigration
func Migrate(){
	db := connectDB()
	defer db.Close()

	db.Debug().AutoMigrate(&User{}, &Account{})

	createAccounts()
}

