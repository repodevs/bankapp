package useraccounts

import (
	"github.com/repodevs/bankapp/helpers"
	"github.com/repodevs/bankapp/interfaces"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB().Debug()
	defer db.Close()

	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
}
