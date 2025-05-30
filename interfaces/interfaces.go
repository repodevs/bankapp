package interfaces

import "github.com/jinzhu/gorm"

// User struct for user
type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

// Account struct for account
type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

// Transaction struct for transations
type Transaction struct {
	gorm.Model
	From 		uint
	To 			uint
	Amount 	int
}

// ResponseAccount used for struct Response from DB
type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

// ResponseUser used for struct Response from DB
type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}

// Validation used for data validation
type Validation struct {
	Value string
	Valid string
}

// Login  for
type Login struct {
	Username string
	Password string
}

// Register used for struct Register
type Register struct {
	Username string
	Email    string
	Password string
}

// TransactionBody struct for transaction
type TransactionBody struct {
	UserID uint
	From uint
	To uint
	Amount int
}

// ErrResponse for
type ErrResponse struct {
	Message string
}
