package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(pass string) (string, error) {
	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("Error on field %s, consition %s")
	}

	return string(hashedPassword), err
}
