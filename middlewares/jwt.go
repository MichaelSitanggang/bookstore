package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 1. header
// 2. payload
// 3. singnature

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	AccountID int    `json:"accountID"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJwt(accountID int, role string) (string, error) {
	ExpiredTime := time.Now().Add(24 * time.Hour)
	claims := Claims{ // payload
		AccountID: accountID,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ExpiredTime),
			Issuer:    "BookStore",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // header
	tokenString, err := token.SignedString(JwtKey)             // singnature
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
