package handlers

import (
	"net/http"
	"strings"
	"time"

	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func GetSetupStatus(c *gin.Context) {
	// Check if ANY users exist — if not, this is a fresh install
	var userCount int64
	database.DB.Model(&models.SystemUser{}).Count(&userCount)
	if userCount == 0 {
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
	c.JSON(http.StatusOK, gin.H{
		"isFirstRun":       false,
		"isSetupCompleted": config.Value == "true",
	})
}

// CreateAdmin bootstraps the very first admin account on a fresh installation.
// This endpoint only works when ZERO users exist in the system.
func CreateAdmin(c *gin.Context) {
	// Safety: refuse if any user already exists
	var userCount int64
	database.DB.Model(&models.SystemUser{}).Count(&userCount)
	if userCount > 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin account already exists. Use the login page."})
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
	req.Email = strings.TrimSpace(req.Email)
	req.EmployeeID = strings.TrimSpace(req.EmployeeID)
	req.Department = strings.TrimSpace(req.Department)
	req.CompanyName = strings.TrimSpace(req.CompanyName)

	if len(strings.TrimSpace(req.Password)) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters"})
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

	// Issue JWT immediately so user is auto-logged in
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "super-secret-tasm-key"
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  admin.ID,
		"role": admin.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token":   tokenString,
		"user":    admin,
		"message": "Admin account created. Please complete the setup wizard.",
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
			if i > 0 {
				cats += ","
			}
			cats += cat
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
	database.DB.Model(&models.SystemConfig{}).Where("key = ?", "is_setup_completed").Update("value", "true")

	c.JSON(http.StatusOK, gin.H{"message": "Setup completed successfully"})
}
