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

func ComparePassword(pass string, hashPass string) bool {
	bytePassword := []byte(pass)
	byteHashPassword := []byte(hashPass)

	err := bcrypt.CompareHashAndPassword(byteHashPassword, bytePassword)

	if err == nil {
		return true
	} else {
		return false
	}
}
