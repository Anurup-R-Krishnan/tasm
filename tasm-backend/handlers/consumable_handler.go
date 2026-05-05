package handlers

import (
	"net/http"
	"strings"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetConsumables(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var consumables []models.Consumable
	if err := db.Order("name asc").Find(&consumables).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load consumables"})
		return
	}

	c.JSON(http.StatusOK, consumables)
}

func GetConsumableByID(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var consumable models.Consumable
	if err := db.First(&consumable, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Consumable not found"})
		return
	}

	c.JSON(http.StatusOK, consumable)
}

func CreateConsumable(c *gin.Context) {
	var consumable models.Consumable
	if err := c.ShouldBindJSON(&consumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consumable.Name = strings.TrimSpace(consumable.Name)
	consumable.Category = strings.TrimSpace(consumable.Category)
	consumable.Location = strings.TrimSpace(consumable.Location)

	switch {
	case consumable.Name == "":
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	case consumable.Category == "":
		c.JSON(http.StatusBadRequest, gin.H{"error": "category is required"})
		return
	case consumable.Location == "":
		c.JSON(http.StatusBadRequest, gin.H{"error": "location is required"})
		return
	case consumable.CurrentStock < 0:
		c.JSON(http.StatusBadRequest, gin.H{"error": "currentStock cannot be negative"})
		return
	case consumable.ReorderLevel < 0:
		c.JSON(http.StatusBadRequest, gin.H{"error": "reorderLevel cannot be negative"})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&consumable).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create consumable"})
		return
	}

	c.JSON(http.StatusCreated, consumable)
}

func UpdateConsumable(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.Consumable
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	var payload models.Consumable
	if !bindJSON(c, &payload) {
		return
	}

	if payload.Name != "" {
		item.Name = trimSpace(payload.Name)
	}
	if payload.Category != "" {
		item.Category = trimSpace(payload.Category)
	}
	if payload.Location != "" {
		item.Location = trimSpace(payload.Location)
	}
	if payload.CurrentStock != 0 {
		item.CurrentStock = payload.CurrentStock
	}
	if payload.ReorderLevel != 0 {
		item.ReorderLevel = payload.ReorderLevel
	}

	if !requireNonEmpty(c, "name", item.Name) ||
		!requireNonEmpty(c, "category", item.Category) ||
		!requireNonEmpty(c, "location", item.Location) {
		return
	}
	if !requireIntPositiveOrZero(c, "currentStock", item.CurrentStock) ||
		!requireIntPositiveOrZero(c, "reorderLevel", item.ReorderLevel) {
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteConsumable(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.Consumable{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
