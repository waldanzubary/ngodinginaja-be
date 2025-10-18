package middleware

import (
	"ngodinginaja-be/config"
	"ngodinginaja-be/models"
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return  func (c*gin.Context)  {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)


		token, err := jwt.Parse(tokenString, func (Token *jwt.Token) (interface{}, error)  {
			return []byte(os.Getenv("JWT_SECRET")), nil
			
		})

			if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token error"})
			c.Abort()
			return
		}

			var user models.User
		config.DB.First(&user, claims["user_id"])

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	
		
	}
}