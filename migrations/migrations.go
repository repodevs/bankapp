package migrations

import (
	"github.com/repodevs/bankapp/helpers"
	"github.com/repodevs/bankapp/interfaces"
)

// createAccounts used for create dummy account
func createAccounts() {
	db := helpers.ConnectDB().Debug()

	defer db.Close()

	users := &[2]interfaces.User{
		{Username: "Anton", Email: "Anton@bank.app"},
		{Username: "Bayu", Email: "Bayu@bank.app"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "daily account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

// Migrate create automigration
func Migrate(){
	db := helpers.ConnectDB().Debug()
	defer db.Close()

	User := &interfaces.User{}
	Account := &interfaces.Account{}

	db.AutoMigrate(&User, &Account)

	createAccounts()
}

