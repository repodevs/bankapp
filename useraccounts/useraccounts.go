package useraccounts

import (
	"fmt"

	"github.com/repodevs/bankapp/helpers"
	"github.com/repodevs/bankapp/interfaces"
	"github.com/repodevs/bankapp/transactions"
)

func getAccount(id uint) *interfaces.Account {
	db := helpers.ConnectDB().Debug()
	defer db.Close()

	account := &interfaces.Account{}
	if db.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}

	return account
}

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	db := helpers.ConnectDB().Debug()
	defer db.Close()

	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	// Update Balance of the account
	db.Where("id = ?", id).First(&account)
	account.Balance = uint(amount)
	db.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)

	return responseAcc
}


// Transaction used for doing transaction
func Transaction(userID uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	// format userID as string
	// refactor it later!
	userIDString := fmt.Sprint(userID)
	isValid := helpers.ValidateToken(userIDString, jwt)

	if !isValid {
		return map[string]interface{}{"message": "Token is not valid!"}
	}

	fromAccount := getAccount(from)
	toAccount := getAccount(to)

	if fromAccount == nil || toAccount == nil {
		return map[string]interface{}{"message": "Account not found"}
	} else if fromAccount.UserID != userID {
		return map[string]interface{}{"message": "You are not owner of the account!"}
	} else if int(fromAccount.Balance) < amount {
		return map[string]interface{}{"message": "Account balance is too small"}
	}

	updatedAccount := updateAccount(from, int(fromAccount.Balance) - amount)
	updateAccount(to, int(toAccount.Balance) + amount)


	transactions.CreateTransaction(from, to, amount)

	var response = map[string]interface{}{"message": "ok"}
	response["data"] = updatedAccount
	return response
}