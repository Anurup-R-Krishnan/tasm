package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetAlerts(c *gin.Context) {
	var alerts []models.SystemAlert
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("id desc").Find(&alerts)

	if len(alerts) == 0 {
		alerts = []models.SystemAlert{
			{Title: "Warranty Expired", Message: "Main Generator B (GEN-B-002) warranty has expired. Immediate renewal required.", Type: "Critical", Source: "Asset Management", IsRead: false},
			{Title: "Maintenance Due", Message: "HVAC Unit 4 (HVAC-004) scheduled maintenance is due in 3 days.", Type: "Warning", Source: "Maintenance", IsRead: false},
			{Title: "Stock Low", Message: "Server Rack Screws inventory dropping below threshold (current: 45, min: 50).", Type: "Info", Source: "Inventory", IsRead: false},
		}
		for _, a := range alerts {
			database.DB.Create(&a)
		}
		database.DB.Order("id desc").Find(&alerts)
	}

	c.JSON(http.StatusOK, alerts)
}
