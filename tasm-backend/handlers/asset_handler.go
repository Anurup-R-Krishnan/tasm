package handlers

import (
	"net/http"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
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

func GetAssetByID(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	c.JSON(http.StatusOK, asset)
}

func UpdateAsset(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
		return
	}

	c.JSON(http.StatusOK, asset)
}

func DeleteAsset(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.Asset{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete asset"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset deleted successfully"})
}
