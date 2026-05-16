package handlers

import (
	"net/http"
	"strings"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var locations []models.Location
	query := db.Order("name asc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if locType := strings.TrimSpace(c.Query("type")); locType != "" {
		query = query.Where("type = ?", locType)
	}
	if err := query.Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load locations"})
		return
	}

	c.JSON(http.StatusOK, locations)
}

func GetLocationByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var location models.Location
	if err := db.First(&location, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}

func CreateLocation(c *gin.Context) {
	var loc models.Location
	if !bindJSON(c, &loc) {
		return
	}

	loc.Name = trimSpace(loc.Name)
	loc.Type = trimSpace(loc.Type)
	loc.Address = trimSpace(loc.Address)
	loc.Status = trimSpace(loc.Status)

	if !requireNonEmpty(c, "name", loc.Name) ||
		!requireNonEmpty(c, "type", loc.Type) {
		return
	}
	if loc.Capacity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capacity cannot be negative"})
		return
	}
	if loc.Status == "" {
		loc.Status = "Active"
	}
	if !validateStatus(c, "status", loc.Status, []string{"Active", "Inactive"}) {
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	var existing models.Location
	if err := db.Where("name = ?", loc.Name).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "location with this name already exists"})
		return
	}

	if err := db.Create(&loc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create location"})
		return
	}

	c.JSON(http.StatusCreated, loc)
}

func UpdateLocation(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.Location
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type locationUpdateRequest struct {
		Name     *string `json:"name"`
		Type     *string `json:"type"`
		Address  *string `json:"address"`
		Capacity *int    `json:"capacity"`
		Status   *string `json:"status"`
	}

	var payload locationUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.Name != nil {
		item.Name = trimSpace(*payload.Name)
	}
	if payload.Type != nil {
		item.Type = trimSpace(*payload.Type)
	}
	if payload.Address != nil {
		item.Address = trimSpace(*payload.Address)
	}
	if payload.Capacity != nil {
		item.Capacity = *payload.Capacity
	}
	if payload.Status != nil {
		item.Status = trimSpace(*payload.Status)
	}

	if !requireNonEmpty(c, "name", item.Name) ||
		!requireNonEmpty(c, "type", item.Type) {
		return
	}
	if item.Capacity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capacity cannot be negative"})
		return
	}
	if item.Status == "" {
		item.Status = "Active"
	}
	if !validateStatus(c, "status", item.Status, []string{"Active", "Inactive"}) {
		return
	}

	if payload.Name != nil {
		var existing models.Location
		if err := db.Where("name = ? AND id <> ?", item.Name, item.ID).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "location with this name already exists"})
			return
		}
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteLocation(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}
	deleteEntity(c, db, &models.Location{}, id)
}
