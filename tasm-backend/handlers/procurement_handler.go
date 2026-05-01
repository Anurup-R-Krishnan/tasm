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
	
	database.DB.Order("id desc").Find(&requests)

	if len(requests) == 0 {
		requests = []models.ProcurementRequest{
			{Title: "New Employee Laptops (x10)", Status: "Pending Approval", Priority: "High", EstimatedValue: 15000, RequestorName: "Alice Smith", Department: "IT", PONumber: "PO-2023-001"},
			{Title: "Standing Desks (x5)", Status: "PO Raised", Priority: "Medium", EstimatedValue: 2500, RequestorName: "Bob Johnson", Department: "HR", PONumber: "PO-2023-002"},
			{Title: "Office Supplies Q3", Status: "Received", Priority: "Low", EstimatedValue: 500, RequestorName: "Carol Williams", Department: "Admin", PONumber: "PO-2023-003"},
		}
		for _, req := range requests {
			database.DB.Create(&req)
		}
		database.DB.Order("id desc").Find(&requests)
	}

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
