package handlers

import (
	"net/http"

	"tasm-backend/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func requireDB(c *gin.Context) (*gorm.DB, bool) {
	if database.DB == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not connected"})
		return nil, false
	}

	return database.DB, true
}
