package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"tasm-backend/database"
	"tasm-backend/models"
	"tasm-backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string            `json:"token"`
	User  models.SystemUser `json:"user"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	var user models.SystemUser
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	if strings.TrimSpace(user.Status) != "Active" {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is inactive"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	})

	tokenString, err := token.SignedString(utils.GetJWTSecret())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Update last login
	user.LastLogin = time.Now()
	database.DB.Save(&user)

	c.JSON(http.StatusOK, LoginResponse{
		Token: tokenString,
		User:  user,
	})
}

func GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.SystemUser
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type RegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	EmployeeId string `json:"employeeId"`
	Department string `json:"department"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Normalize email
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	// Flaw 1: Block Register during First Run to prevent admin lockout
	var setupConfig models.SystemConfig
	err := database.DB.Where("key = ?", "is_setup_completed").First(&setupConfig).Error
	if err != nil {
		// If setup config doesn't even exist, the system might be in first-run state
		var adminCount int64
		database.DB.Model(&models.SystemUser{}).Count(&adminCount)
		if adminCount == 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "System is not initialized. Please run the setup wizard first."})
			return
		}
	}

	// Flaw 2: Password Complexity
	if len(req.Password) < 8 || !strings.ContainsAny(req.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") || !strings.ContainsAny(req.Password, "abcdefghijklmnopqrstuvwxyz") || !strings.ContainsAny(req.Password, "0123456789") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters and contain uppercase, lowercase, and numbers."})
		return
	}

	// Flaw 10: Duplicate Email Handling (Check first, then rely on DB constraint)
	var count int64
	database.DB.Model(&models.SystemUser{}).Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.SystemUser{
		Name:         req.Name,
		Email:        req.Email,
		EmployeeID:   req.EmployeeId,
		Department:   req.Department,
		Role:         "Employee",
		PasswordHash: string(hashedPassword),
		Status:       "Active",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unique") {
			c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Flaw 8: Update UserRole count
	database.DB.Model(&models.UserRole{}).Where("role_name = ?", "Employee").UpdateColumn("users_count", gorm.Expr("users_count + ?", 1))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(utils.GetJWTSecret())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	user.LastLogin = time.Now()
	database.DB.Save(&user)

	c.JSON(http.StatusCreated, LoginResponse{
		Token: tokenString,
		User:  user,
	})
}
