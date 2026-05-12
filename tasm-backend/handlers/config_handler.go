package handlers

import (
	"net/http"
	"strings"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetSystemConfig(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var configs []models.SystemConfig
	if err := db.Find(&configs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load configuration"})
		return
	}

	payload := map[string]string{}
	for _, cfg := range configs {
		payload[cfg.Key] = cfg.Value
	}

	c.JSON(http.StatusOK, payload)
}

type updateConfigRequest struct {
	Configs map[string]string `json:"configs"`
}

func UpdateSystemConfig(c *gin.Context) {
	var req updateConfigRequest
	if !bindJSON(c, &req) {
		return
	}

	if len(req.Configs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "configs is required"})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	for key, value := range req.Configs {
		cleanKey := strings.TrimSpace(key)
		if cleanKey == "" {
			continue
		}
		cleanValue := strings.TrimSpace(value)
		db.Where(models.SystemConfig{Key: cleanKey}).
			Assign(models.SystemConfig{Value: cleanValue}).
			FirstOrCreate(&models.SystemConfig{})
	}

	c.JSON(http.StatusOK, req.Configs)
}
