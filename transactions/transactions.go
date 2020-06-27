package transactions

import (
	"github.com/repodevs/bankapp/helpers"
	"github.com/repodevs/bankapp/interfaces"
)

// CreateTransaction used to create a transaction between accounts
func CreateTransaction(from uint, to uint, amount int) {
	db := helpers.ConnectDB().Debug()
	defer db.Close()

	transaction := &interfaces.Transaction{From: from, To: to, Amount: amount}
	db.Create(&transaction)
}