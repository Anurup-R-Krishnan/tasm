package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetProcurements(c *gin.Context) {
	var requests []models.ProcurementRequest
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Find(&requests)
	c.JSON(http.StatusOK, requests)
}

func CreateProcurement(c *gin.Context) {
	var request models.ProcurementRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Create(&request)
	c.JSON(http.StatusCreated, request)
}
