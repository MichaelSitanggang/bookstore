package middlewares

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		HeaderAuth := c.GetHeader("Authorization")
		if HeaderAuth == "" {
			c.JSON(401, gin.H{"message": "Header kosong"})
			c.Abort()
			return
		}
		splits := strings.Split(HeaderAuth, " ")
		if len(splits) != 2 || splits[0] != "Bearer" {
			c.JSON(401, gin.H{"message": "bearer error"})
			c.Abort()
			return
		}
		tokenString := splits[1]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) { // singnature
			return JwtKey, nil

		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"message": "token error"})
			c.Abort()
			return
		}
		c.Set("accountID", claims.AccountID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

//6108503494
