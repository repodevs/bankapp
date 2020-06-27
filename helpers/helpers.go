package helpers

import (
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/repodevs/bankapp/interfaces"
	"golang.org/x/crypto/bcrypt"
)

// HandleErr used for handling error
func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// HashAndSalt used for hasing password using bcrypt
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

// ConnectDB used for connection to DB
func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "postgres://postgres:123456@localhost:5432/go_bankapp?sslmode=disable")
	HandleErr(err)
	return db
}

// Validation used to verify data
// verify email pattern
// verify username pattern
// verify password at least 5 chars long
func Validation(values []interfaces.Validation) bool {
	username := regexp.MustCompile("^([A-Za-z0-9]{5,})+$")
	email := regexp.MustCompile("^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$")

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}

	return true
}
