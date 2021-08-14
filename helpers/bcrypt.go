package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func CompareHash(password string, dbPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(dbPass))
	return err
}
