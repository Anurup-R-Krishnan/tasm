package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var systemAdminRoles = []string{"System Admin"}

func RequireRoles(allowed ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User role not found"})
			return
		}

		role, ok := roleValue.(string)
		if !ok || strings.TrimSpace(role) == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User role not found"})
			return
		}

		if isRoleAllowed(role, allowed) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
	}
}

func isRoleAllowed(role string, allowed []string) bool {
	for _, adminRole := range systemAdminRoles {
		if role == adminRole {
			return true
		}
	}

	for _, option := range allowed {
		if role == option {
			return true
		}
	}

	return false
}
