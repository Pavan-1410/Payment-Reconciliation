package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int `json:"user_id"`
	IS_Admin bool `json:"is_admin"`
	jwt.RegisteredClaims		// This embeds the JWT library's standard claims struct into your custom struct.
}

func GenerateToken(userID int, isAdmin bool) (string, error) {

	claims := Claims{
		UserID: userID,
		IS_Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}