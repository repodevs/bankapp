package helpers

import "golang.org/x/crypto/bcrypt"

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
