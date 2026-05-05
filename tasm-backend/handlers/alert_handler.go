package handlers

import (
	"net/http"
	"strings"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAlerts(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var alerts []models.SystemAlert
	query := db.Order("created_at desc")
	if alertType := strings.TrimSpace(c.Query("type")); alertType != "" {
		query = query.Where("type = ?", alertType)
	}
	if c.Query("unread") == "true" {
		query = query.Where("is_read = false")
	}
	if err := query.Find(&alerts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load alerts"})
		return
	}

	c.JSON(http.StatusOK, alerts)
}

func GetAlertByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var alert models.SystemAlert
	if err := db.First(&alert, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	c.JSON(http.StatusOK, alert)
}

func UpdateAlert(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.SystemAlert
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type alertUpdateRequest struct {
		Title   *string `json:"title"`
		Message *string `json:"message"`
		Type    *string `json:"type"`
		Source  *string `json:"source"`
		IsRead  *bool   `json:"isRead"`
	}

	var payload alertUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.Title != nil {
		item.Title = trimSpace(*payload.Title)
	}
	if payload.Message != nil {
		item.Message = trimSpace(*payload.Message)
	}
	if payload.Type != nil {
		item.Type = trimSpace(*payload.Type)
	}
	if payload.Source != nil {
		item.Source = trimSpace(*payload.Source)
	}
	if payload.IsRead != nil {
		item.IsRead = *payload.IsRead
	}

	if !requireNonEmpty(c, "title", item.Title) ||
		!requireNonEmpty(c, "message", item.Message) ||
		!requireNonEmpty(c, "type", item.Type) {
		return
	}
	if !validateStatus(c, "type", item.Type, []string{"Critical", "Warning", "Info"}) {
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteAlert(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.SystemAlert{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
