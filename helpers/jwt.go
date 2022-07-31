package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	ID    int
	Email string
}

func GenerateToken(Id int, email string) (string, error) {
	claims := &AuthClaims{
		jwt.StandardClaims{
			Issuer:    "Big-Backend-Golang",
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
		Id,
		email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("thisissecretkey"))

	return t, err
}

func DecodeToken(token string) int {
	claim := &AuthClaims{}
	jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte("thisissecretkey"), nil
	})
	fmt.Println(claim, "ini data claim")
	return claim.ID
}
