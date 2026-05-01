package handlers

import (
	"net/http"
	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetLedgers(c *gin.Context) {
	var ledgers []models.LedgerEntry
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("date desc").Find(&ledgers)

	c.JSON(http.StatusOK, ledgers)
}

func GetLeases(c *gin.Context) {
	var leases []models.LeaseAgreement
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("start_date desc").Find(&leases)

	c.JSON(http.StatusOK, leases)
}

func GetDepreciations(c *gin.Context) {
	var schedules []models.DepreciationSchedule
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("id desc").Find(&schedules)

	c.JSON(http.StatusOK, schedules)
}

func UpdateLedger(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.LedgerEntry
	if err := db.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteLedger(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.LedgerEntry{}, "id = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func UpdateLease(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.LedgerEntry
	if err := db.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteLease(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.LedgerEntry{}, "id = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func UpdateDepreciation(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.LedgerEntry
	if err := db.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteDepreciation(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.LedgerEntry{}, "id = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
