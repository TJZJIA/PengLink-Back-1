package auth

import (
	"PengLink-Back-1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户信息
		user := &models.User{}
		if !Authorize(user, requiredRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
