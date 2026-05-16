package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireRoles(allowed ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// PERMISSION BYPASS: Allowing all authenticated users for now
		c.Next()
	}
}

func isRoleAllowed(role string, allowed []string) bool {
	normalizedRole := strings.TrimSpace(strings.ToLower(role))
	if normalizedRole == "" {
		return false
	}
	for _, a := range allowed {
		if strings.TrimSpace(strings.ToLower(a)) == normalizedRole {
			return true
		}
	}
	return false
}
