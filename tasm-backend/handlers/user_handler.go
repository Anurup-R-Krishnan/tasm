package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"tasm-backend/database"
	"tasm-backend/models"
)

// GetUsers returns a list of system users
func GetUsers(c *gin.Context) {
	var users []models.SystemUser
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Seed data if empty
	if len(users) == 0 {
		users = []models.SystemUser{
			{
				EmployeeID: "EMP-001",
				Name:       "Rahul Menon",
				Email:      "rahul.m@technopark.org",
				Department: "IT Operations",
				Role:       "System Admin",
				Status:     "Active",
				LastLogin:  time.Now(),
			},
			{
				EmployeeID: "EMP-042",
				Name:       "Priya Suresh",
				Email:      "priya.s@technopark.org",
				Department: "Finance",
				Role:       "Finance Manager",
				Status:     "Active",
				LastLogin:  time.Now().Add(-2 * time.Hour),
			},
			{
				EmployeeID: "EMP-088",
				Name:       "Anil Kumar",
				Email:      "anil.k@technopark.org",
				Department: "Facilities",
				Role:       "Facilities Head",
				Status:     "Inactive",
				LastLogin:  time.Now().Add(-48 * time.Hour),
			},
		}
		for _, u := range users {
			database.DB.Create(&u)
		}
		database.DB.Find(&users)
	}

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
	if len(roles) == 0 {
		roles = []models.UserRole{
			{
				RoleName:    "System Admin",
				Description: "Full access to all system modules and configuration.",
				UsersCount:  3,
			},
			{
				RoleName:    "Finance Manager",
				Description: "Access to ledgers, depreciation, and lease agreements.",
				UsersCount:  5,
			},
			{
				RoleName:    "Facilities Head",
				Description: "Access to maintenance, work orders, and asset registry.",
				UsersCount:  8,
			},
			{
				RoleName:    "Employee",
				Description: "Self-service portal access for checking out assets.",
				UsersCount:  245,
			},
		}
		for _, r := range roles {
			database.DB.Create(&r)
		}
		database.DB.Find(&roles)
	}

	c.JSON(http.StatusOK, roles)
}
