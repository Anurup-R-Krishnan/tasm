package handlers

import (
	"net/http"
	"strings"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAudits(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var audits []models.AuditSession
	query := db.Order("start_date desc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Find(&audits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load audits"})
		return
	}

	c.JSON(http.StatusOK, audits)
}

func GetAuditByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var audit models.AuditSession
	if err := db.First(&audit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Audit session not found"})
		return
	}

	c.JSON(http.StatusOK, audit)
}

func CreateAudit(c *gin.Context) {
	var audit models.AuditSession
	if !bindJSON(c, &audit) {
		return
	}

	audit.Title = trimSpace(audit.Title)
	audit.Status = trimSpace(audit.Status)
	audit.AuditorName = trimSpace(audit.AuditorName)

	if !requireNonEmpty(c, "title", audit.Title) ||
		!requireNonEmpty(c, "auditorName", audit.AuditorName) {
		return
	}
	if audit.StartDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate is required"})
		return
	}
	if audit.EndDate != nil && audit.EndDate.Before(audit.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endDate cannot be before startDate"})
		return
	}
	if audit.Status == "" {
		audit.Status = "Active"
	}
	if !validateStatus(c, "status", audit.Status, []string{"Active", "Completed"}) {
		return
	}
	if audit.TotalAssets < 0 || audit.ScannedAssets < 0 || audit.DiscrepancyCount < 0 || audit.Progress < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "counts cannot be negative"})
		return
	}
	if audit.TotalAssets > 0 && audit.ScannedAssets > audit.TotalAssets {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scannedAssets cannot exceed totalAssets"})
		return
	}
	if audit.Progress > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "progress cannot exceed 100"})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&audit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create audit"})
		return
	}

	c.JSON(http.StatusCreated, audit)
}

func GetDiscrepancies(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var discs []models.AuditDiscrepancy
	if err := db.Order("id desc").Find(&discs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load discrepancies"})
		return
	}

	c.JSON(http.StatusOK, discs)
}

func GetDiscrepancyByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var discrepancy models.AuditDiscrepancy
	if err := db.First(&discrepancy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Discrepancy not found"})
		return
	}

	c.JSON(http.StatusOK, discrepancy)
}

func UpdateAudit(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.AuditSession
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	if !updateAuditFromPayload(c, &item) {
		return
	}

	if !validateAuditFields(c, &item) {
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func updateAuditFromPayload(c *gin.Context, item *models.AuditSession) bool {
	var payload struct {
		Title            *string `json:"title"`
		StartDate        *string `json:"startDate"`
		EndDate          *string `json:"endDate"`
		Status           *string `json:"status"`
		TotalAssets      *int    `json:"totalAssets"`
		ScannedAssets    *int    `json:"scannedAssets"`
		DiscrepancyCount *int    `json:"discrepancyCount"`
		AuditorName      *string `json:"auditorName"`
		Progress         *int    `json:"progress"`
	}

	if !bindJSON(c, &payload) {
		return false
	}

	if payload.Title != nil {
		item.Title = trimSpace(*payload.Title)
	}
	if payload.StartDate != nil {
		parsed, err := parseTime(*payload.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "startDate must be a valid datetime"})
			return false
		}
		item.StartDate = parsed
	}
	if payload.EndDate != nil {
		parsed, err := parseDatePointer(*payload.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "endDate must be a valid datetime"})
			return false
		}
		item.EndDate = parsed
	}
	if payload.Status != nil {
		item.Status = trimSpace(*payload.Status)
	}
	if payload.TotalAssets != nil {
		item.TotalAssets = *payload.TotalAssets
	}
	if payload.ScannedAssets != nil {
		item.ScannedAssets = *payload.ScannedAssets
	}
	if payload.DiscrepancyCount != nil {
		item.DiscrepancyCount = *payload.DiscrepancyCount
	}
	if payload.AuditorName != nil {
		item.AuditorName = trimSpace(*payload.AuditorName)
	}
	if payload.Progress != nil {
		item.Progress = *payload.Progress
	}
	return true
}

func validateAuditFields(c *gin.Context, item *models.AuditSession) bool {
	if !requireNonEmpty(c, "title", item.Title) ||
		!requireNonEmpty(c, "auditorName", item.AuditorName) {
		return false
	}
	if item.StartDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate is required"})
		return false
	}
	if item.EndDate != nil && item.EndDate.Before(item.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endDate cannot be before startDate"})
		return false
	}
	if item.Status == "" {
		item.Status = "Active"
	}
	if !validateStatus(c, "status", item.Status, []string{"Active", "Completed"}) {
		return false
	}
	if item.TotalAssets < 0 || item.ScannedAssets < 0 || item.DiscrepancyCount < 0 || item.Progress < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "counts cannot be negative"})
		return false
	}
	if item.TotalAssets > 0 && item.ScannedAssets > item.TotalAssets {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scannedAssets cannot exceed totalAssets"})
		return false
	}
	if item.Progress > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "progress cannot exceed 100"})
		return false
	}
	return true
}

func DeleteAudit(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.AuditSession{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func UpdateDiscrepancy(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.AuditDiscrepancy
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	var payload models.AuditDiscrepancy
	if !bindJSON(c, &payload) {
		return
	}

	if payload.AssetName != "" {
		item.AssetName = trimSpace(payload.AssetName)
	}
	if payload.AssetTag != "" {
		item.AssetTag = trimSpace(payload.AssetTag)
	}
	if payload.IssueType != "" {
		item.IssueType = trimSpace(payload.IssueType)
	}
	if payload.LastKnownLocation != "" {
		item.LastKnownLocation = trimSpace(payload.LastKnownLocation)
	}
	if payload.ScannedLocation != "" {
		item.ScannedLocation = trimSpace(payload.ScannedLocation)
	}
	if payload.RecommendedAction != "" {
		item.RecommendedAction = trimSpace(payload.RecommendedAction)
	}

	if !requireNonEmpty(c, "assetTag", item.AssetTag) ||
		!requireNonEmpty(c, "issueType", item.IssueType) {
		return
	}
	if !validateStatus(c, "issueType", item.IssueType, []string{"Missing", "Location Mismatch", "Unregistered"}) {
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteDiscrepancy(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.AuditDiscrepancy{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
