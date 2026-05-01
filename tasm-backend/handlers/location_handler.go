package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetLocations(c *gin.Context) {
	var locations []models.Location
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Order("id desc").Find(&locations)

	if len(locations) == 0 {
		locations = []models.Location{
			{Name: "Headquarters", Type: "Office", Address: "123 Main St, New York, NY", Capacity: 500, Status: "Active"},
			{Name: "Data Center Alpha", Type: "Data Center", Address: "45 Tech Park, Austin, TX", Capacity: 50, Status: "Active"},
			{Name: "West Coast Hub", Type: "Office", Address: "900 Market St, San Francisco, CA", Capacity: 200, Status: "Active"},
		}
		for _, l := range locations {
			database.DB.Create(&l)
		}
		database.DB.Order("id desc").Find(&locations)
	}

	c.JSON(http.StatusOK, locations)
}

func CreateLocation(c *gin.Context) {
	var loc models.Location
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Create(&loc)
	c.JSON(http.StatusCreated, loc)
}
