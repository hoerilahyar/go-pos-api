package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	appErr "gopos/pkg/errors"
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
			return "", time.Time{}, appErr.Get(appErr.ErrMarshalJson, err)
		}
		dataStr = string(jsonBytes)
	}

	encryptedData, err := Encrypt(dataStr)
	if err != nil {
		return "", time.Time{}, appErr.Get(appErr.ErrEncrypt, err)
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
	return signedToken, expireAt, appErr.Get(appErr.ErrTokenSignatureInvalid, err)
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
			// return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, appErr.Get(appErr.ErrSignatureMethod, nil)
		}
		return jwtSecret, nil
	})

	if err != nil {
		// Check for specific JWT errors
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, appErr.Get(appErr.ErrTokenMalformed, err)
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, appErr.Get(appErr.ErrTokenNotYetValid, err)
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, appErr.Get(appErr.ErrTokenSignatureInvalid, err)
		}

		fmt.Printf("Error during token parsing: %v\n", err.Error()) // More detailed error
		return nil, appErr.Get(appErr.ErrTokenInvalid, err)
	}

	if !token.Valid {
		return nil, appErr.Get(appErr.ErrTokenInvalidAfterParsing, nil)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, appErr.Get(appErr.ErrTokenClaimsParsingFailed, nil)
	}

	// Your original expiration check (ParseWithClaims usually handles this, but good to have)
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, appErr.Get(appErr.ErrTokenExpiredSecondary, nil)
	}

	return claims, nil
}

func JwtAuthInfo(c *gin.Context) (interface{}, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, appErr.Get(appErr.ErrAuthHeaderMissing, nil)
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, appErr.Get(appErr.ErrAuthHeaderInvalidFormat, nil)
	}

	tokenStr := parts[1]

	claims, err := ValidateToken(tokenStr)
	if err != nil {
		return nil, appErr.Get(appErr.ErrTokenInvalid, err)
	}

	decryptedData, err := Decrypt(claims.Data)
	if err != nil {
		return nil, appErr.Get(appErr.ErrDecrypt, err)
	}

	return CheckIfJSON(decryptedData), nil
}
