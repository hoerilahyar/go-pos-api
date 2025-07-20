package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

type CustomClaims struct {
	Data string `json:"data"`
	jwt.RegisteredClaims
}

func GenerateToken(data interface{}) (string, time.Time, error) {
	envSecret := os.Getenv("JWT_SECRET")
	if len(envSecret) == 0 {
		jwtSecret = []byte("fallback-secret")
	} else {
		jwtSecret = []byte(strings.TrimSpace(envSecret))
	}

	expireAt := time.Now().Add(72 * time.Hour)

	var dataStr string
	switch v := data.(type) {
	case string:
		dataStr = v
	default:
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return "", time.Time{}, fmt.Errorf("failed to marshal data: %w", err)
		}
		dataStr = string(jsonBytes)
	}

	encryptedData, err := Encrypt(dataStr)
	if err != nil {
		return "", time.Time{}, err
	}

	claims := CustomClaims{
		Data: encryptedData,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	return signedToken, expireAt, err
}

func ValidateToken(tokenStr string) (*CustomClaims, error) {
	// Add this line to see the secret used for validation
	envSecret := os.Getenv("JWT_SECRET")
	if len(envSecret) == 0 {
		jwtSecret = []byte("fallback-secret")
	} else {
		jwtSecret = []byte(strings.TrimSpace(envSecret))
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is what you expect (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		fmt.Printf("Error during token parsing: %v\n", err) // More detailed error
		// Check for specific JWT errors
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("token malformed")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("token expired or not yet valid") // Combine these
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, errors.New("invalid token signature") // Explicitly catch this
		}
		return nil, errors.New("invalid token: " + err.Error()) // General error with details
	}

	if !token.Valid {
		return nil, errors.New("token is not valid after parsing (shouldn't happen with specific error checks)")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("could not parse claims") // If type assertion fails
	}

	// Your original expiration check (ParseWithClaims usually handles this, but good to have)
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired (secondary check)")
	}

	return claims, nil
}

func JwtAuthInfo(c *gin.Context) (interface{}, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("Authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("Invalid Authorization header format")
	}

	tokenStr := parts[1]

	claims, err := ValidateToken(tokenStr)
	if err != nil {
		return nil, err
	}

	decryptedData, err := Decrypt(claims.Data)
	if err != nil {
		return nil, err
	}

	return CheckIfJSON(decryptedData), nil
}
