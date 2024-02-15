package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "suchsecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
