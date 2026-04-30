package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tasm-backend/models"
)

func GetAssets(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var assets []models.Asset
	if err := db.Order("created_at desc").Find(&assets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load assets"})
		return
	}

	c.JSON(http.StatusOK, assets)
}

func CreateAsset(c *gin.Context) {
	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create asset"})
		return
	}

	c.JSON(http.StatusCreated, asset)
}
