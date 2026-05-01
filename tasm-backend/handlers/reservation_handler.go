package handlers

import (
	"net/http"

	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetReservations(c *gin.Context) {
	var reservations []models.Reservation
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("id desc").Find(&reservations)

	c.JSON(http.StatusOK, reservations)
}

func UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.Reservation
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

func DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.Reservation{}, "id = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
