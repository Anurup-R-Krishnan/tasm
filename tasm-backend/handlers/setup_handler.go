package handlers

import (
	"net/http"
	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetSetupStatus(c *gin.Context) {
	var config models.SystemConfig
	if err := database.DB.Where("key = ?", "is_setup_completed").First(&config).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"isSetupCompleted": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"isSetupCompleted": config.Value == "true"})
}

func CompleteSetup(c *gin.Context) {
	var req struct {
		CompanyName string            `json:"companyName" binding:"required"`
		Currency    string            `json:"currency" binding:"required"`
		Locations   []models.Location `json:"locations"`
		Categories  []string          `json:"categories"`
		Features    map[string]bool   `json:"features"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save configs
	configs := []models.SystemConfig{
		{Key: "company_name", Value: req.CompanyName},
		{Key: "currency", Value: req.Currency},
	}

	for k, v := range req.Features {
		val := "false"
		if v {
			val = "true"
		}
		configs = append(configs, models.SystemConfig{Key: "feature_" + k, Value: val})
	}

	for _, cfg := range configs {
		database.DB.Where(models.SystemConfig{Key: cfg.Key}).Assign(models.SystemConfig{Value: cfg.Value}).FirstOrCreate(&models.SystemConfig{})
	}

	// Create locations
	for _, loc := range req.Locations {
		loc.Status = "Active"
		database.DB.Create(&loc)
	}

	database.DB.Model(&models.SystemConfig{}).Where("key = ?", "is_setup_completed").Update("value", "true")

	c.JSON(http.StatusOK, gin.H{"message": "Setup completed successfully"})
}
