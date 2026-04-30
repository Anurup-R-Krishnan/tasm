package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tasm-backend/models"
)

func GetContracts(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var contracts []models.MaintenanceContract
	if err := db.Order("start_date desc").Find(&contracts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load contracts"})
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func CreateContract(c *gin.Context) {
	var contract models.MaintenanceContract
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&contract).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contract"})
		return
	}

	c.JSON(http.StatusCreated, contract)
}

func GetWorkOrders(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var workOrders []models.WorkOrder
	if err := db.Order("target_date asc").Find(&workOrders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load work orders"})
		return
	}

	c.JSON(http.StatusOK, workOrders)
}

func CreateWorkOrder(c *gin.Context) {
	var workOrder models.WorkOrder
	if err := c.ShouldBindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&workOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create work order"})
		return
	}

	c.JSON(http.StatusCreated, workOrder)
}
