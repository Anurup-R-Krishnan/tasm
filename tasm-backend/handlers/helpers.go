package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func requireDB(c *gin.Context) (*gorm.DB, bool) {
	if database.DB == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not connected"})
		return nil, false
	}

	return database.DB, true
}

func parseIDParam(c *gin.Context) (uint, bool) {
	value := strings.TrimSpace(c.Param("id"))
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return 0, false
	}

	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}

	return uint(id), true
}

func bindJSON(c *gin.Context, target interface{}) bool {
	if err := c.ShouldBindJSON(target); err != nil {
		mapValidationErrors(c, err)
		return false
	}

	return true
}

func trimSpace(value string) string {
	return strings.TrimSpace(value)
}

func requireNonEmpty(c *gin.Context, field string, value string) bool {
	if strings.TrimSpace(value) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": field + " is required"})
		return false
	}

	return true
}

func requirePositiveOrZero(c *gin.Context, field string, value float64) bool {
	if value < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": field + " cannot be negative"})
		return false
	}

	return true
}

func requireIntPositiveOrZero(c *gin.Context, field string, value int) bool {
	if value < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": field + " cannot be negative"})
		return false
	}

	return true
}

func validateStatus(c *gin.Context, field string, value string, allowed []string) bool {
	for _, option := range allowed {
		if value == option {
			return true
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": field + " must be one of: " + strings.Join(allowed, ", "),
	})
	return false
}

var dateLayouts = []string{
	time.RFC3339,
	"2006-01-02T15:04:05",
	"2006-01-02",
}

func parseTime(value string) (time.Time, error) {
	if strings.TrimSpace(value) == "" {
		return time.Time{}, errors.New("time value is required")
	}

	for _, layout := range dateLayouts {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			return parsed, nil
		}
	}

	return time.Time{}, errors.New("invalid time format")
}

func parseDatePointer(value string) (*time.Time, error) {
	if strings.TrimSpace(value) == "" {
		return nil, nil
	}

	for _, layout := range dateLayouts {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			return &parsed, nil
		}
	}

	return nil, errors.New("invalid date format")
}

func mapValidationErrors(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
}

func requirePasswordPolicy(c *gin.Context, field string, value string) bool {
	trimmed := strings.TrimSpace(value)
	if len(trimmed) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": field + " must be at least 8 characters"})
		return false
	}

	var hasUpper bool
	var hasLower bool
	var hasNumber bool
	var hasSpecial bool
	for _, char := range trimmed {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case strings.ContainsRune("!\"#$%&'()*+,-./:;=?@[\\]^_`{|}~", char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		c.JSON(http.StatusBadRequest, gin.H{"error": field + " must include upper, lower, number, and special character"})
		return false
	}

	return true
}

func deleteEntity(c *gin.Context, db *gorm.DB, model interface{}, id uint) {
	if err := db.Delete(model, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

// getActorFromContext extracts the authenticated user's ID and name from the Gin context.
func getActorFromContext(c *gin.Context, db *gorm.DB) (uint, string) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, "System"
	}
	actorID := userID.(uint)
	var actor models.SystemUser
	if err := db.First(&actor, actorID).Error; err == nil {
		return actorID, actor.Name
	}
	return actorID, "System"
}
