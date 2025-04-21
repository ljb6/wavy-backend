package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwt_secret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		// search for the token
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}

		// validating
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwt_secret, nil
		})
		
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("userID", claims["userID"])

		c.Next()
	}
}