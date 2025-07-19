package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // Load from env ideally

type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func GenerateToken(userID uint) (string, time.Time, error) {
	expireAt := time.Now().Add(72 * time.Hour) // token valid 3 days

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expireAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)
	return t, expireAt, err
}
