package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"tasm-backend/models"
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
