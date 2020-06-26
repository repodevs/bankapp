package helpers

import (
	"github.com/jinzhu/gorm"
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