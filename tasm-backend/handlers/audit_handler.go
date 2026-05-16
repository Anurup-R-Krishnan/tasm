package handlers

import (
	"net/http"
	"strings"
	"time"

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

	// Auto-set totalAssets from DB count if not provided
	if audit.TotalAssets == 0 {
		var count int64
		db.Model(&models.Asset{}).Where("status IN ?", []string{"In Stock", "Checked Out", "Reserved"}).Count(&count)
		audit.TotalAssets = int(count)
	}

	if err := db.Create(&audit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create audit"})
		return
	}

	c.JSON(http.StatusCreated, audit)
}

// ScanAssetInAudit registers a scanned asset tag against an active audit session.
// It updates scanned count and progress; if there's a discrepancy it records it.
func ScanAssetInAudit(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var session models.AuditSession
	if err := db.First(&session, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Audit session not found"})
		return
	}
	if session.Status != "Active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Audit session is not active"})
		return
	}

	type scanPayload struct {
		TagID           string `json:"tagId"`
		ScannedLocation string `json:"scannedLocation"`
	}
	var payload scanPayload
	if !bindJSON(c, &payload) {
		return
	}
	payload.TagID = trimSpace(payload.TagID)
	if payload.TagID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tagId is required"})
		return
	}

	// Check if already scanned in this session
	var existingScan models.AuditScan
	err := db.Where("audit_session_id = ? AND asset_tag = ?", session.ID, payload.TagID).First(&existingScan).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "already_scanned",
			"message": "Asset already scanned in this session",
		})
		return
	}

	// Look up the asset by tag
	var asset models.Asset
	result := db.Where("tag_id = ?", payload.TagID).First(&asset)

	if result.Error != nil {
		// Asset not found — log as unregistered discrepancy
		disc := models.AuditDiscrepancy{
			AuditSessionID:    session.ID,
			AssetTag:          payload.TagID,
			AssetName:         "Unknown",
			IssueType:         "Unregistered",
			ScannedLocation:   payload.ScannedLocation,
			LastKnownLocation: "N/A",
			RecommendedAction: "Register as new asset or investigate",
			Status:            "Open",
		}
		db.Create(&disc)

		// Record the scan even if not found in registry (to prevent double-counting discrepancies)
		db.Create(&models.AuditScan{
			AuditSessionID: session.ID,
			AssetTag:       payload.TagID,
			ScannedAt:      time.Now(),
		})

		// Increment discrepancy count
		session.DiscrepancyCount++
		db.Save(&session)

		c.JSON(http.StatusOK, gin.H{
			"result":      "unregistered",
			"message":     "Asset tag not found in system",
			"discrepancy": disc,
		})
		return
	}

	// Asset found — check for location mismatch
	scannedLoc := trimSpace(payload.ScannedLocation)
	locationMismatch := scannedLoc != "" && scannedLoc != asset.Location

	// Record the scan
	db.Create(&models.AuditScan{
		AuditSessionID: session.ID,
		AssetTag:       asset.TagID,
		ScannedAt:      time.Now(),
	})

	// Update scanned count and progress
	session.ScannedAssets++
	if session.TotalAssets > 0 {
		session.Progress = (session.ScannedAssets * 100) / session.TotalAssets
	}

	if locationMismatch {
		// Check if a discrepancy for this asset in this audit already exists
		var existingDisc models.AuditDiscrepancy
		discErr := db.Where("audit_session_id = ? AND asset_tag = ? AND status = ?", session.ID, asset.TagID, "Open").First(&existingDisc).Error
		if discErr != nil {
			disc := models.AuditDiscrepancy{
				AuditSessionID:    session.ID,
				AssetTag:          asset.TagID,
				AssetName:         asset.Name,
				IssueType:         "Location Mismatch",
				LastKnownLocation: asset.Location,
				ScannedLocation:   scannedLoc,
				RecommendedAction: "Verify asset location and update record",
				Status:            "Open",
			}
			db.Create(&disc)
			session.DiscrepancyCount++
		}
	}

	db.Save(&session)

	c.JSON(http.StatusOK, gin.H{
		"result":        "found",
		"asset":         asset,
		"locationMatch": !locationMismatch,
		"progress":      session.Progress,
	})
}

// ResolveDiscrepancy marks a discrepancy as resolved and optionally updates the related asset
func ResolveDiscrepancy(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var disc models.AuditDiscrepancy
	if err := db.First(&disc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Discrepancy not found"})
		return
	}
	if disc.Status == "Resolved" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Discrepancy is already resolved"})
		return
	}

	type resolvePayload struct {
		Action      string `json:"action"`      // confirm_location, mark_lost, register, update_location, dismiss
		NewLocation string `json:"newLocation"` // for update_location action
		Notes       string `json:"notes"`
	}
	var payload resolvePayload
	if !bindJSON(c, &payload) {
		return
	}
	payload.Action = trimSpace(payload.Action)
	if payload.Action == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "action is required"})
		return
	}

	validActions := []string{"confirm_location", "mark_lost", "register", "update_location", "dismiss"}
	if !validateStatus(c, "action", payload.Action, validActions) {
		return
	}

	// Get actor name from auth context
	actorName := "System"
	if v, exists := c.Get("userID"); exists {
		actorID := v.(uint)
		var actor models.SystemUser
		if err := db.First(&actor, actorID).Error; err == nil {
			actorName = actor.Name
		}
	}

	// Apply the action
	switch payload.Action {
	case "confirm_location":
		// Asset location is correct as scanned — update DB to match scan
		if disc.AssetTag != "" && disc.ScannedLocation != "" {
			db.Model(&models.Asset{}).Where("tag_id = ?", disc.AssetTag).Update("location", disc.ScannedLocation)
		}
		disc.Resolution = "Location confirmed and updated to: " + disc.ScannedLocation

	case "mark_lost":
		if disc.AssetTag != "" {
			db.Model(&models.Asset{}).Where("tag_id = ?", disc.AssetTag).Update("status", "Retired")
		}
		disc.Resolution = "Asset marked as lost/retired"

	case "register":
		// For unregistered assets, we just dismiss — user should register via AddNewAsset
		disc.Resolution = "Flagged for registration as new asset"

	case "update_location":
		newLoc := trimSpace(payload.NewLocation)
		if newLoc == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "newLocation is required for update_location action"})
			return
		}
		if disc.AssetTag != "" {
			db.Model(&models.Asset{}).Where("tag_id = ?", disc.AssetTag).Update("location", newLoc)
		}
		disc.Resolution = "Location updated to: " + newLoc

	case "dismiss":
		disc.Resolution = "Dismissed: " + payload.Notes
	}

	disc.Status = "Resolved"
	disc.ResolvedBy = actorName
	disc.UpdatedAt = time.Now()

	if err := db.Save(&disc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resolve discrepancy"})
		return
	}

	// Recalculate discrepancy count for the audit session
	if disc.AuditSessionID > 0 {
		var openCount int64
		db.Model(&models.AuditDiscrepancy{}).Where("audit_session_id = ? AND status = ?", disc.AuditSessionID, "Open").Count(&openCount)
		db.Model(&models.AuditSession{}).Where("id = ?", disc.AuditSessionID).Update("discrepancy_count", openCount)
	}

	c.JSON(http.StatusOK, disc)
}

// CreateDiscrepancy creates a new discrepancy record
func CreateDiscrepancy(c *gin.Context) {
	var disc models.AuditDiscrepancy
	if !bindJSON(c, &disc) {
		return
	}
	disc.AssetTag = trimSpace(disc.AssetTag)
	disc.IssueType = trimSpace(disc.IssueType)
	disc.AssetName = trimSpace(disc.AssetName)

	if !requireNonEmpty(c, "assetTag", disc.AssetTag) ||
		!requireNonEmpty(c, "issueType", disc.IssueType) {
		return
	}
	if !validateStatus(c, "issueType", disc.IssueType, []string{"Missing", "Location Mismatch", "Unregistered"}) {
		return
	}
	if disc.Status == "" {
		disc.Status = "Open"
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}
	if err := db.Create(&disc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create discrepancy"})
		return
	}

	// Update audit discrepancy count if linked
	if disc.AuditSessionID > 0 {
		var openCount int64
		db.Model(&models.AuditDiscrepancy{}).Where("audit_session_id = ? AND status = ?", disc.AuditSessionID, "Open").Count(&openCount)
		db.Model(&models.AuditSession{}).Where("id = ?", disc.AuditSessionID).Update("discrepancy_count", openCount)
	}

	c.JSON(http.StatusCreated, disc)
}

func GetDiscrepancies(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var discs []models.AuditDiscrepancy
	query := db.Order("id desc")

	if auditID := c.Query("auditSessionId"); auditID != "" {
		query = query.Where("audit_session_id = ?", auditID)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if issueType := c.Query("issueType"); issueType != "" {
		query = query.Where("issue_type = ?", issueType)
	}

	if err := query.Find(&discs).Error; err != nil {
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
	deleteEntity(c, db, &models.AuditSession{}, id)
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
	if payload.AuditSessionID > 0 {
		item.AuditSessionID = payload.AuditSessionID
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
	deleteEntity(c, db, &models.AuditDiscrepancy{}, id)
}
