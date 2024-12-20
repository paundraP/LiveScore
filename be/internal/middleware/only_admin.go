package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OnlyAdmin(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access"})
			c.Abort()
			return
		}
		c.Next()
	}
}
