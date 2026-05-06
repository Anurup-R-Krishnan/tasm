package handlers

import (
	"net/http"
	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SeedDatabase(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	// Wipe all tables
	database.DB.Exec("TRUNCATE TABLE assets, audit_sessions, audit_discrepancies, consumables, maintenance_contracts, work_orders, procurement_requests, ledger_entries, lease_agreements, depreciation_schedules, software_licenses, system_users, user_roles, reservations, locations, system_alerts, system_configs RESTART IDENTITY CASCADE")

	// Create roles
	roles := []models.UserRole{
		{RoleName: "System Admin", Description: "Full system access", UsersCount: 1},
		{RoleName: "Finance Manager", Description: "Financial oversight", UsersCount: 0},
		{RoleName: "Facilities Head", Description: "Operations and maintenance", UsersCount: 0},
		{RoleName: "Employee", Description: "Standard access", UsersCount: 0},
	}
	for _, role := range roles {
		database.DB.Create(&role)
	}

	// Create initial admin
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := models.SystemUser{
		EmployeeID:   "ADMIN-001",
		Name:         "System Admin",
		Email:        "admin@tasm.com",
		PasswordHash: string(hashedPassword),
		Role:         "System Admin",
		Status:       "Active",
		Department:   "IT Services",
	}
	database.DB.Create(&admin)

	// Set setup pending
	config := models.SystemConfig{
		Key:   "is_setup_completed",
		Value: "false",
	}
	database.DB.Create(&config)

	c.JSON(http.StatusOK, gin.H{"message": "Database wiped and initial admin created. Credentials: admin@tasm.com / admin123"})
}
