package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	UserId int64 `json:"user_id"`
}

func ValidateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if auth != "123456" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()

	}
}
