package middleware

import (
	"errors"
	"fmt"
	"gopos/pkg/response"
	"gopos/pkg/utils"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, errors.New("Authorization header missing"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Error(c, errors.New("Authorization header format must be Bearer {token}"))
			c.Abort()
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Unexpected signing method")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, errors.New("Invalid token"))
			c.Abort()
			return
		}

		// You may set claims or user data here

		claims, err := utils.JwtAuthInfo(c)
		if err != nil {
			response.Error(c, errors.New("Invalid token"))
			c.Abort()
			return
		}

		parsedClaims, ok := claims.(map[string]interface{})
		if !ok {
			response.Error(c, errors.New("Invalid token claims format"))
			c.Abort()
			return
		}

		userIDFloat, ok := parsedClaims["user_id"].(float64)
		if !ok {
			response.Error(c, errors.New("Invalid user_id type"))
			c.Abort()
			return
		}

		userIDStr := fmt.Sprintf("%.0f", userIDFloat)
		c.Set("user_id", userIDStr)

		c.Next()
	}
}
