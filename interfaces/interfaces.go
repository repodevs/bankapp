package interfaces

import "github.com/jinzhu/gorm"

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

// ResponseAccount used for struct Response from DB
type ResponseAccount struct {
	ID uint
	Name string
	Balance uint
}

// ResponseUser used for struct Response from DB
type ResponseUser struct {
	ID uint
	Username string
	Email string
	Accounts []ResponseAccount
}

