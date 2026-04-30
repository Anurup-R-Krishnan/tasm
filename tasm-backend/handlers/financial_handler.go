package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetLedgers(c *gin.Context) {
	var ledgers []models.LedgerEntry
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Find(&ledgers)
	c.JSON(http.StatusOK, ledgers)
}

func GetLeases(c *gin.Context) {
	var leases []models.LeaseAgreement
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Find(&leases)
	c.JSON(http.StatusOK, leases)
}

func GetDepreciations(c *gin.Context) {
	var schedules []models.DepreciationSchedule
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Find(&schedules)
	c.JSON(http.StatusOK, schedules)
}
