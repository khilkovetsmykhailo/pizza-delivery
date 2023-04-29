package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secret_key = "123456"

func HashPassword(password string) string {
	bytePass := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	return string(hashedPassword)
}

func ComparePassword(dbPass, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(password)) == nil
}

func GenerateToken(username, role string) string {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
		"role":     role,
		"username": username,
	}

	var signingKey = []byte(secret_key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(signingKey)
	return t
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret_key), nil
	})
}
