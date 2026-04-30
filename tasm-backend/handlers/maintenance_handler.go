package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetContracts(c *gin.Context) {
	var contracts []models.MaintenanceContract
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Find(&contracts)
	c.JSON(http.StatusOK, contracts)
}

func CreateContract(c *gin.Context) {
	var contract models.MaintenanceContract
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Create(&contract)
	c.JSON(http.StatusCreated, contract)
}

func GetWorkOrders(c *gin.Context) {
	var workOrders []models.WorkOrder
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Find(&workOrders)
	c.JSON(http.StatusOK, workOrders)
}

func CreateWorkOrder(c *gin.Context) {
	var workOrder models.WorkOrder
	if err := c.ShouldBindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Create(&workOrder)
	c.JSON(http.StatusCreated, workOrder)
}
