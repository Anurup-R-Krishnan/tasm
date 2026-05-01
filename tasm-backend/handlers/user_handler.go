package handlers

import (
	"net/http"

	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

// GetUsers returns a list of system users
func GetUsers(c *gin.Context) {
	var users []models.SystemUser
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Seed data if empty

	c.JSON(http.StatusOK, users)
}

// GetRoles returns a list of user roles
func GetRoles(c *gin.Context) {
	var roles []models.UserRole
	if err := database.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch roles"})
		return
	}

	// Seed data if empty

	c.JSON(http.StatusOK, roles)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.SystemUser
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

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.SystemUser{}, "id = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.SystemUser
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

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.SystemUser{}, "id = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
