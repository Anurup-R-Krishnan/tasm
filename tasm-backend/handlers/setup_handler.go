package handlers

import (
	"net/http"
	"strings"
	"time"

	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetSetupStatus(c *gin.Context) {
	// Flaw 5: Admin Resurrection Prevention
	// Don't just check userCount, check if the system config was ever initialized
	var configCount int64
	database.DB.Model(&models.SystemConfig{}).Where("key = ?", "is_setup_completed").Count(&configCount)

	if configCount == 0 {
		c.JSON(http.StatusOK, gin.H{
			"isFirstRun":       true,
			"isSetupCompleted": false,
		})
		return
	}

	var config models.SystemConfig
	if err := database.DB.Where("key = ?", "is_setup_completed").First(&config).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"isFirstRun":       false,
			"isSetupCompleted": false,
		})
		return
	}
	var companyName models.SystemConfig
	database.DB.Where("key = ?", "company_name").First(&companyName)

	c.JSON(http.StatusOK, gin.H{
		"isFirstRun":       false,
		"isSetupCompleted": config.Value == "true",
		"companyName":      companyName.Value,
	})
}

// CreateAdmin bootstraps the very first admin account on a fresh installation.
// This endpoint only works when ZERO users exist in the system.
func CreateAdmin(c *gin.Context) {
	// Safety: refuse if the system has already been initialized
	var configCount int64
	database.DB.Model(&models.SystemConfig{}).Where("key = ?", "is_setup_completed").Count(&configCount)
	if configCount > 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin account already exists or system is initialized. Use the login page."})
		return
	}

	var req struct {
		Name        string `json:"name" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required"`
		EmployeeID  string `json:"employeeId"`
		Department  string `json:"department"`
		CompanyName string `json:"companyName"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email)) // Flaw 4: Normalize email
	req.EmployeeID = strings.TrimSpace(req.EmployeeID)
	req.Department = strings.TrimSpace(req.Department)
	req.CompanyName = strings.TrimSpace(req.CompanyName)

	// Flaw 2: Password Complexity
	if len(req.Password) < 8 || !strings.ContainsAny(req.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") || !strings.ContainsAny(req.Password, "abcdefghijklmnopqrstuvwxyz") || !strings.ContainsAny(req.Password, "0123456789") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters and contain uppercase, lowercase, and numbers."})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	employeeID := req.EmployeeID
	if employeeID == "" {
		employeeID = "ADMIN-001"
	}
	department := req.Department
	if department == "" {
		department = "IT Services"
	}

	// Create default roles
	defaultRoles := []models.UserRole{
		{RoleName: "System Admin", Description: "Full system access — can manage users, roles, and all configuration", UsersCount: 1},
		{RoleName: "Finance Manager", Description: "Manage financial records, ledgers, leases, and depreciation", UsersCount: 0},
		{RoleName: "Facilities Head", Description: "Manage assets, locations, maintenance, and audit sessions", UsersCount: 0},
		{RoleName: "Employee", Description: "Standard access — can view assets and submit self-service requests", UsersCount: 0},
	}
	for _, role := range defaultRoles {
		database.DB.Where(models.UserRole{RoleName: role.RoleName}).FirstOrCreate(&role)
	}

	admin := models.SystemUser{
		EmployeeID:   employeeID,
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         "System Admin",
		Status:       "Active",
		Department:   department,
		LastLogin:    time.Now(),
	}
	if err := database.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin account"})
		return
	}

	// Save company name if provided
	if req.CompanyName != "" {
		database.DB.Where(models.SystemConfig{Key: "company_name"}).
			Assign(models.SystemConfig{Value: req.CompanyName}).
			FirstOrCreate(&models.SystemConfig{})
	}

	// Seed the is_setup_completed key as false so wizard runs next
	database.DB.Where(models.SystemConfig{Key: "is_setup_completed"}).
		Assign(models.SystemConfig{Value: "false"}).
		FirstOrCreate(&models.SystemConfig{})

	// Do not auto-issue a JWT here. Require explicit sign-in to continue.
	c.JSON(http.StatusCreated, gin.H{
		"user":    admin,
		"message": "Admin account created. Please sign in to continue.",
	})
}

func CompleteSetup(c *gin.Context) {
	var req struct {
		CompanyName string            `json:"companyName" binding:"required"`
		Currency    string            `json:"currency" binding:"required"`
		Locations   []models.Location `json:"locations"`
		Categories  []string          `json:"categories"`
		Fields      map[string]bool   `json:"fields"`
		Features    map[string]bool   `json:"features"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Flaw 3: Replay Attack in CompleteSetup
	var setupConfig models.SystemConfig
	database.DB.Where("key = ?", "is_setup_completed").First(&setupConfig)
	if setupConfig.Value == "true" {
		c.JSON(http.StatusConflict, gin.H{"error": "Setup has already been completed."})
		return
	}

	// Save core configs
	configs := []models.SystemConfig{
		{Key: "company_name", Value: req.CompanyName},
		{Key: "currency", Value: req.Currency},
	}

	// Save features
	for k, v := range req.Features {
		val := "false"
		if v {
			val = "true"
		}
		configs = append(configs, models.SystemConfig{Key: "feature_" + k, Value: val})
	}

	// Save field preferences
	for k, v := range req.Fields {
		val := "false"
		if v {
			val = "true"
		}
		configs = append(configs, models.SystemConfig{Key: "field_" + k, Value: val})
	}

	// Save categories as comma-separated string
	if len(req.Categories) > 0 {
		cats := ""
		for i, cat := range req.Categories {
			// Flaw 7: Unsanitized Categories
			cleanCat := strings.ReplaceAll(cat, ",", " ")
			if i > 0 {
				cats += ","
			}
			cats += cleanCat
		}
		configs = append(configs, models.SystemConfig{Key: "asset_categories", Value: cats})
	}

	for _, cfg := range configs {
		database.DB.Where(models.SystemConfig{Key: cfg.Key}).Assign(models.SystemConfig{Value: cfg.Value}).FirstOrCreate(&models.SystemConfig{})
	}

	// Create locations
	for _, loc := range req.Locations {
		if loc.Name == "" {
			continue
		}
		loc.Status = "Active"
		database.DB.Create(&loc)
	}

	// Mark setup as completed
	database.DB.Where(models.SystemConfig{Key: "is_setup_completed"}).
		Assign(models.SystemConfig{Value: "true"}).
		FirstOrCreate(&models.SystemConfig{})

	c.JSON(http.StatusOK, gin.H{"message": "Setup completed successfully"})
}
