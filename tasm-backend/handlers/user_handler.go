package handlers

import (
	"net/http"
	"strings"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers returns a list of system users
func GetUsers(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var users []models.SystemUser
	query := db.Order("created_at desc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if role := strings.TrimSpace(c.Query("role")); role != "" {
		query = query.Where("role = ?", role)
	}
	if department := strings.TrimSpace(c.Query("department")); department != "" {
		query = query.Where("department = ?", department)
	}
	if err := query.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var user models.SystemUser
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type createUserRequest struct {
	EmployeeID string `json:"employeeId"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Department string `json:"department"`
	Role       string `json:"role"`
	Status     string `json:"status"`
}

type updateUserRequest struct {
	EmployeeID *string `json:"employeeId"`
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Department *string `json:"department"`
	Role       *string `json:"role"`
	Status     *string `json:"status"`
	Password   *string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var payload createUserRequest
	if !bindJSON(c, &payload) {
		return
	}

	payload.EmployeeID = trimSpace(payload.EmployeeID)
	payload.Name = trimSpace(payload.Name)
	payload.Email = strings.ToLower(trimSpace(payload.Email))
	payload.Department = trimSpace(payload.Department)
	payload.Role = trimSpace(payload.Role)
	payload.Status = trimSpace(payload.Status)

	if !requireNonEmpty(c, "employeeId", payload.EmployeeID) ||
		!requireNonEmpty(c, "name", payload.Name) ||
		!requireNonEmpty(c, "email", payload.Email) ||
		!requireNonEmpty(c, "password", payload.Password) {
		return
	}
	if payload.Status == "" {
		payload.Status = "Active"
	}
	if !validateStatus(c, "status", payload.Status, []string{"Active", "Inactive"}) {
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	var existing models.SystemUser
	if err := db.Where("email = ?", payload.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user with this email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to secure password"})
		return
	}

	user := models.SystemUser{
		EmployeeID:   payload.EmployeeID,
		Name:         payload.Name,
		Email:        payload.Email,
		PasswordHash: string(hashedPassword),
		Department:   payload.Department,
		Role:         payload.Role,
		Status:       payload.Status,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetRoles returns a list of user roles
func GetRoles(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var roles []models.UserRole
	query := db.Order("role_name asc")
	if err := query.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch roles"})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func GetRoleByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var role models.UserRole
	if err := db.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func CreateRole(c *gin.Context) {
	var role models.UserRole
	if !bindJSON(c, &role) {
		return
	}

	role.RoleName = trimSpace(role.RoleName)
	role.Description = trimSpace(role.Description)
	if !requireNonEmpty(c, "roleName", role.RoleName) {
		return
	}
	if role.UsersCount < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usersCount cannot be negative"})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	var existing models.UserRole
	if err := db.Where("role_name = ?", role.RoleName).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "role with this name already exists"})
		return
	}

	if err := db.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func UpdateUser(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.SystemUser
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	var payload updateUserRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.EmployeeID != nil {
		item.EmployeeID = trimSpace(*payload.EmployeeID)
	}
	if payload.Name != nil {
		item.Name = trimSpace(*payload.Name)
	}
	if payload.Email != nil {
		item.Email = strings.ToLower(trimSpace(*payload.Email))
	}
	if payload.Department != nil {
		item.Department = trimSpace(*payload.Department)
	}
	if payload.Role != nil {
		item.Role = trimSpace(*payload.Role)
	}
	if payload.Status != nil {
		item.Status = trimSpace(*payload.Status)
	}
	if payload.Password != nil && strings.TrimSpace(*payload.Password) != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*payload.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to secure password"})
			return
		}
		item.PasswordHash = string(hashedPassword)
	}

	if !requireNonEmpty(c, "employeeId", item.EmployeeID) ||
		!requireNonEmpty(c, "name", item.Name) ||
		!requireNonEmpty(c, "email", item.Email) {
		return
	}
	if item.Status == "" {
		item.Status = "Active"
	}
	if !validateStatus(c, "status", item.Status, []string{"Active", "Inactive"}) {
		return
	}

	if payload.Email != nil {
		var existing models.SystemUser
		if err := db.Where("email = ? AND id <> ?", item.Email, item.ID).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "user with this email already exists"})
			return
		}
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteUser(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.SystemUser{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func UpdateRole(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.UserRole
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type roleUpdateRequest struct {
		RoleName    *string `json:"roleName"`
		Description *string `json:"description"`
		UsersCount  *int    `json:"usersCount"`
	}

	var payload roleUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.RoleName != nil {
		item.RoleName = trimSpace(*payload.RoleName)
	}
	if payload.Description != nil {
		item.Description = trimSpace(*payload.Description)
	}
	if payload.UsersCount != nil {
		item.UsersCount = *payload.UsersCount
	}

	if !requireNonEmpty(c, "roleName", item.RoleName) {
		return
	}
	if item.UsersCount < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usersCount cannot be negative"})
		return
	}

	if payload.RoleName != nil {
		var existing models.UserRole
		if err := db.Where("role_name = ? AND id <> ?", item.RoleName, item.ID).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "role with this name already exists"})
			return
		}
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteRole(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.UserRole{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
