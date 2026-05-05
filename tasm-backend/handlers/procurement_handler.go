package handlers

import (
	"net/http"
	"strings"

	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProcurements(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var requests []models.ProcurementRequest
	query := db.Order("created_at desc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if priority := strings.TrimSpace(c.Query("priority")); priority != "" {
		query = query.Where("priority = ?", priority)
	}
	if department := strings.TrimSpace(c.Query("department")); department != "" {
		query = query.Where("department = ?", department)
	}
	if err := query.Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load procurements"})
		return
	}

	c.JSON(http.StatusOK, requests)
}

func GetProcurementByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	var request models.ProcurementRequest
	if err := database.DB.First(&request, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Procurement request not found"})
		return
	}

	c.JSON(http.StatusOK, request)
}

func CreateProcurement(c *gin.Context) {
	var request models.ProcurementRequest
	if !bindJSON(c, &request) {
		return
	}

	request.Title = trimSpace(request.Title)
	request.Status = trimSpace(request.Status)
	request.Priority = trimSpace(request.Priority)
	request.RequestorName = trimSpace(request.RequestorName)
	request.RequestorInitials = trimSpace(request.RequestorInitials)
	request.Department = trimSpace(request.Department)
	request.PONumber = trimSpace(request.PONumber)

	if !requireNonEmpty(c, "title", request.Title) ||
		!requireNonEmpty(c, "department", request.Department) ||
		!requireNonEmpty(c, "requestorName", request.RequestorName) {
		return
	}
	if request.Priority == "" {
		request.Priority = "Low"
	}
	if !validateStatus(c, "priority", request.Priority, []string{"Low", "Medium", "High"}) {
		return
	}
	if request.Status == "" {
		request.Status = "Draft"
	}
	if !validateStatus(c, "status", request.Status, []string{"Draft", "Pending Approval", "PO Raised", "Shipping", "Received"}) {
		return
	}
	if !requirePositiveOrZero(c, "estimatedValue", request.EstimatedValue) ||
		!requirePositiveOrZero(c, "actualValue", request.ActualValue) {
		return
	}
	if request.RequestorInitials == "" {
		parts := strings.Fields(request.RequestorName)
		initials := ""
		for _, part := range parts {
			initials += strings.ToUpper(string(part[0]))
			if len(initials) == 2 {
				break
			}
		}
		request.RequestorInitials = initials
	}

	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	if err := database.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create procurement"})
		return
	}

	c.JSON(http.StatusCreated, request)
}

func UpdateProcurement(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.ProcurementRequest
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type procurementUpdateRequest struct {
		Title             *string  `json:"title"`
		Status            *string  `json:"status"`
		Priority          *string  `json:"priority"`
		EstimatedValue    *float64 `json:"estimatedValue"`
		ActualValue       *float64 `json:"actualValue"`
		RequestorName     *string  `json:"requestorName"`
		RequestorInitials *string  `json:"requestorInitials"`
		Department        *string  `json:"department"`
		PONumber          *string  `json:"poNumber"`
	}

	var payload procurementUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.Title != nil {
		item.Title = trimSpace(*payload.Title)
	}
	if payload.Status != nil {
		item.Status = trimSpace(*payload.Status)
	}
	if payload.Priority != nil {
		item.Priority = trimSpace(*payload.Priority)
	}
	if payload.EstimatedValue != nil {
		item.EstimatedValue = *payload.EstimatedValue
	}
	if payload.ActualValue != nil {
		item.ActualValue = *payload.ActualValue
	}
	if payload.RequestorName != nil {
		item.RequestorName = trimSpace(*payload.RequestorName)
	}
	if payload.RequestorInitials != nil {
		item.RequestorInitials = trimSpace(*payload.RequestorInitials)
	}
	if payload.Department != nil {
		item.Department = trimSpace(*payload.Department)
	}
	if payload.PONumber != nil {
		item.PONumber = trimSpace(*payload.PONumber)
	}

	if !requireNonEmpty(c, "title", item.Title) ||
		!requireNonEmpty(c, "department", item.Department) ||
		!requireNonEmpty(c, "requestorName", item.RequestorName) {
		return
	}
	if item.Priority == "" {
		item.Priority = "Low"
	}
	if !validateStatus(c, "priority", item.Priority, []string{"Low", "Medium", "High"}) {
		return
	}
	if item.Status == "" {
		item.Status = "Draft"
	}
	if !validateStatus(c, "status", item.Status, []string{"Draft", "Pending Approval", "PO Raised", "Shipping", "Received"}) {
		return
	}
	if !requirePositiveOrZero(c, "estimatedValue", item.EstimatedValue) ||
		!requirePositiveOrZero(c, "actualValue", item.ActualValue) {
		return
	}
	if item.RequestorInitials == "" {
		parts := strings.Fields(item.RequestorName)
		initials := ""
		for _, part := range parts {
			initials += strings.ToUpper(string(part[0]))
			if len(initials) == 2 {
				break
			}
		}
		item.RequestorInitials = initials
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteProcurement(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.ProcurementRequest{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
