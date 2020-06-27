package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/repodevs/bankapp/helpers"
	"github.com/repodevs/bankapp/interfaces"
	"golang.org/x/crypto/bcrypt"
)

// Login used for checking credential user and log in it
func Login(username string, pass string) map[string]interface{} {
	// Check if paramaters is valid
	isValid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})

	if isValid {
		// Connect to DB
		db := helpers.ConnectDB().Debug()
		defer db.Close()

		// Check user in DB
		user := &interfaces.User{}
		if db.Where("username = ?", username).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		// Check password
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Wrong password"}
		}

		// Find accounts of user
		accounts := []interfaces.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		var response = prepareResponse(user, accounts)

		return response
	}

	return map[string]interface{}{"message": "not valid values"}
}

// Register is used for registering user
func Register(username string, email string, pass string) map[string]interface{} {
	isValid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})

	if isValid {
		// Connect to DB
		db := helpers.ConnectDB().Debug()
		defer db.Close()

		// Create user
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{Username: username, Email: email, Password: generatedPassword}
		db.Create(&user)

		// Create account
		account := &interfaces.Account{Type: "daily account", Name: string(username + "'s" + " account"), Balance: uint(0), UserID: user.ID}
		db.Create(&account)

		// Prepare response
		accounts := []interfaces.ResponseAccount{}
		respAccount := interfaces.ResponseAccount{ID: account.ID, Name: account.Name, Balance: uint(account.Balance)}
		accounts = append(accounts, respAccount)
		var response = prepareResponse(user, accounts)

		return response
	}

	return map[string]interface{}{"message": "not valid values"}
}

func prepareToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	var token = prepareToken(user)
	var response = map[string]interface{}{"message": "ok"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}
