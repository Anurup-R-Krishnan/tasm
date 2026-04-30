package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetAudits(c *gin.Context) {
	var audits []models.AuditSession
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Find(&audits)
	c.JSON(http.StatusOK, audits)
}

func CreateAudit(c *gin.Context) {
	var audit models.AuditSession
	if err := c.ShouldBindJSON(&audit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Create(&audit)
	c.JSON(http.StatusCreated, audit)
}
