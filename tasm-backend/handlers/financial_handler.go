package handlers

import (
	"net/http"
	"strings"
	"time"

	"tasm-backend/database"
	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetLedgers(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var ledgers []models.LedgerEntry
	query := db.Order("date desc")
	if entryType := strings.TrimSpace(c.Query("type")); entryType != "" {
		query = query.Where("type = ?", entryType)
	}
	if category := strings.TrimSpace(c.Query("category")); category != "" {
		query = query.Where("category = ?", category)
	}
	if err := query.Find(&ledgers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load ledgers"})
		return
	}

	c.JSON(http.StatusOK, ledgers)
}

func GetSoftwareLicenses(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var licenses []models.SoftwareLicense
	query := db.Order("renewal_date asc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Find(&licenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load software licenses"})
		return
	}

	c.JSON(http.StatusOK, licenses)
}

func CreateLedger(c *gin.Context) {
	var ledger models.LedgerEntry
	if !bindJSON(c, &ledger) {
		return
	}

	ledger.TransactionID = trimSpace(ledger.TransactionID)
	ledger.Description = trimSpace(ledger.Description)
	ledger.Type = trimSpace(ledger.Type)
	ledger.Category = trimSpace(ledger.Category)

	if !requireNonEmpty(c, "transactionId", ledger.TransactionID) ||
		!requireNonEmpty(c, "description", ledger.Description) ||
		!requireNonEmpty(c, "type", ledger.Type) {
		return
	}
	if ledger.Date.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date is required"})
		return
	}
	if !validateStatus(c, "type", ledger.Type, []string{"Credit", "Debit"}) {
		return
	}
	if !requirePositiveOrZero(c, "amount", ledger.Amount) {
		return
	}
	if ledger.Date.After(time.Now().AddDate(1, 0, 0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date is not realistic"})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&ledger).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ledger entry"})
		return
	}

	c.JSON(http.StatusCreated, ledger)
}

func GetLedgerByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	var ledger models.LedgerEntry
	if err := database.DB.First(&ledger, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ledger entry not found"})
		return
	}

	c.JSON(http.StatusOK, ledger)
}

func GetLeases(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var leases []models.LeaseAgreement
	query := db.Order("start_date desc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if vendor := strings.TrimSpace(c.Query("vendor")); vendor != "" {
		query = query.Where("vendor ILIKE ?", "%"+vendor+"%")
	}
	if err := query.Find(&leases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load leases"})
		return
	}

	c.JSON(http.StatusOK, leases)
}

func CreateLease(c *gin.Context) {
	var lease models.LeaseAgreement
	if !bindJSON(c, &lease) {
		return
	}

	lease.LeaseID = trimSpace(lease.LeaseID)
	lease.Vendor = trimSpace(lease.Vendor)
	lease.Status = trimSpace(lease.Status)

	if !requireNonEmpty(c, "leaseId", lease.LeaseID) ||
		!requireNonEmpty(c, "vendor", lease.Vendor) {
		return
	}
	if lease.StartDate.IsZero() || lease.EndDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate and endDate are required"})
		return
	}
	if lease.EndDate.Before(lease.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endDate cannot be before startDate"})
		return
	}
	if lease.Status == "" {
		lease.Status = "Active"
	}
	if !validateStatus(c, "status", lease.Status, []string{"Active", "Expired"}) {
		return
	}
	if !requirePositiveOrZero(c, "monthlyCost", lease.MonthlyCost) {
		return
	}
	if lease.StartDate.After(time.Now().AddDate(10, 0, 0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate is not realistic"})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&lease).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lease"})
		return
	}

	c.JSON(http.StatusCreated, lease)
}

func GetLeaseByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	var lease models.LeaseAgreement
	if err := database.DB.First(&lease, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lease agreement not found"})
		return
	}

	c.JSON(http.StatusOK, lease)
}

func GetDepreciations(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var schedules []models.DepreciationSchedule
	query := db.Order("id desc")
	if method := strings.TrimSpace(c.Query("method")); method != "" {
		query = query.Where("method = ?", method)
	}
	if err := query.Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load depreciations"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func CreateDepreciation(c *gin.Context) {
	var schedule models.DepreciationSchedule
	if !bindJSON(c, &schedule) {
		return
	}

	schedule.AssetID = trimSpace(schedule.AssetID)
	schedule.AssetName = trimSpace(schedule.AssetName)
	schedule.Method = trimSpace(schedule.Method)

	if !requireNonEmpty(c, "assetId", schedule.AssetID) ||
		!requireNonEmpty(c, "assetName", schedule.AssetName) ||
		!requireNonEmpty(c, "method", schedule.Method) {
		return
	}
	if !validateStatus(c, "method", schedule.Method, []string{"Straight Line", "Declining Balance"}) {
		return
	}
	if !requirePositiveOrZero(c, "purchaseValue", schedule.PurchaseValue) ||
		!requirePositiveOrZero(c, "currentValue", schedule.CurrentValue) {
		return
	}
	if schedule.CurrentValue > schedule.PurchaseValue {
		c.JSON(http.StatusBadRequest, gin.H{"error": "currentValue cannot exceed purchaseValue"})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create depreciation schedule"})
		return
	}

	c.JSON(http.StatusCreated, schedule)
}

func GetDepreciationByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	var schedule models.DepreciationSchedule
	if err := database.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Depreciation schedule not found"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func GetSoftwareLicenseByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	var license models.SoftwareLicense
	if err := database.DB.First(&license, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Software license not found"})
		return
	}

	c.JSON(http.StatusOK, license)
}

func CreateSoftwareLicense(c *gin.Context) {
	var license models.SoftwareLicense
	if !bindJSON(c, &license) {
		return
	}

	license.SoftwareName = trimSpace(license.SoftwareName)
	license.PlanName = trimSpace(license.PlanName)
	license.Status = trimSpace(license.Status)

	if !requireNonEmpty(c, "softwareName", license.SoftwareName) ||
		!requireNonEmpty(c, "planName", license.PlanName) {
		return
	}
	if license.Status == "" {
		license.Status = "Active"
	}
	if !validateStatus(c, "status", license.Status, []string{"Active", "Expiring Soon", "Expired"}) {
		return
	}
	if !requireIntPositiveOrZero(c, "totalSeats", license.TotalSeats) ||
		!requireIntPositiveOrZero(c, "usedSeats", license.UsedSeats) {
		return
	}
	if license.UsedSeats > license.TotalSeats {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usedSeats cannot exceed totalSeats"})
		return
	}
	if license.RenewalDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "renewalDate is required"})
		return
	}
	if !requirePositiveOrZero(c, "annualCost", license.AnnualCost) {
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&license).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create software license"})
		return
	}

	c.JSON(http.StatusCreated, license)
}

func UpdateLedger(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.LedgerEntry
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type ledgerUpdateRequest struct {
		TransactionID *string    `json:"transactionId"`
		Date          *time.Time `json:"date"`
		Description   *string    `json:"description"`
		Amount        *float64   `json:"amount"`
		Type          *string    `json:"type"`
		Category      *string    `json:"category"`
	}

	var payload ledgerUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.TransactionID != nil {
		item.TransactionID = trimSpace(*payload.TransactionID)
	}
	if payload.Date != nil {
		item.Date = *payload.Date
	}
	if payload.Description != nil {
		item.Description = trimSpace(*payload.Description)
	}
	if payload.Amount != nil {
		item.Amount = *payload.Amount
	}
	if payload.Type != nil {
		item.Type = trimSpace(*payload.Type)
	}
	if payload.Category != nil {
		item.Category = trimSpace(*payload.Category)
	}

	if !requireNonEmpty(c, "transactionId", item.TransactionID) ||
		!requireNonEmpty(c, "description", item.Description) ||
		!requireNonEmpty(c, "type", item.Type) {
		return
	}
	if item.Date.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date is required"})
		return
	}
	if !validateStatus(c, "type", item.Type, []string{"Credit", "Debit"}) {
		return
	}
	if !requirePositiveOrZero(c, "amount", item.Amount) {
		return
	}
	if item.Date.After(time.Now().AddDate(1, 0, 0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date is not realistic"})
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteLedger(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.LedgerEntry{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func UpdateLease(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.LeaseAgreement
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type leaseUpdateRequest struct {
		LeaseID     *string    `json:"leaseId"`
		Vendor      *string    `json:"vendor"`
		StartDate   *time.Time `json:"startDate"`
		EndDate     *time.Time `json:"endDate"`
		MonthlyCost *float64   `json:"monthlyCost"`
		Status      *string    `json:"status"`
	}

	var payload leaseUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.LeaseID != nil {
		item.LeaseID = trimSpace(*payload.LeaseID)
	}
	if payload.Vendor != nil {
		item.Vendor = trimSpace(*payload.Vendor)
	}
	if payload.StartDate != nil {
		item.StartDate = *payload.StartDate
	}
	if payload.EndDate != nil {
		item.EndDate = *payload.EndDate
	}
	if payload.MonthlyCost != nil {
		item.MonthlyCost = *payload.MonthlyCost
	}
	if payload.Status != nil {
		item.Status = trimSpace(*payload.Status)
	}

	if !requireNonEmpty(c, "leaseId", item.LeaseID) ||
		!requireNonEmpty(c, "vendor", item.Vendor) {
		return
	}
	if item.StartDate.IsZero() || item.EndDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate and endDate are required"})
		return
	}
	if item.EndDate.Before(item.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endDate cannot be before startDate"})
		return
	}
	if item.Status == "" {
		item.Status = "Active"
	}
	if !validateStatus(c, "status", item.Status, []string{"Active", "Expired"}) {
		return
	}
	if !requirePositiveOrZero(c, "monthlyCost", item.MonthlyCost) {
		return
	}
	if item.StartDate.After(time.Now().AddDate(10, 0, 0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate is not realistic"})
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteLease(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.LeaseAgreement{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func UpdateDepreciation(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.DepreciationSchedule
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type depreciationUpdateRequest struct {
		AssetID       *string  `json:"assetId"`
		AssetName     *string  `json:"assetName"`
		PurchaseValue *float64 `json:"purchaseValue"`
		CurrentValue  *float64 `json:"currentValue"`
		Method        *string  `json:"method"`
	}

	var payload depreciationUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.AssetID != nil {
		item.AssetID = trimSpace(*payload.AssetID)
	}
	if payload.AssetName != nil {
		item.AssetName = trimSpace(*payload.AssetName)
	}
	if payload.PurchaseValue != nil {
		item.PurchaseValue = *payload.PurchaseValue
	}
	if payload.CurrentValue != nil {
		item.CurrentValue = *payload.CurrentValue
	}
	if payload.Method != nil {
		item.Method = trimSpace(*payload.Method)
	}

	if !requireNonEmpty(c, "assetId", item.AssetID) ||
		!requireNonEmpty(c, "assetName", item.AssetName) ||
		!requireNonEmpty(c, "method", item.Method) {
		return
	}
	if !validateStatus(c, "method", item.Method, []string{"Straight Line", "Declining Balance"}) {
		return
	}
	if !requirePositiveOrZero(c, "purchaseValue", item.PurchaseValue) ||
		!requirePositiveOrZero(c, "currentValue", item.CurrentValue) {
		return
	}
	if item.CurrentValue > item.PurchaseValue {
		c.JSON(http.StatusBadRequest, gin.H{"error": "currentValue cannot exceed purchaseValue"})
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func UpdateSoftwareLicense(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.SoftwareLicense
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type licenseUpdateRequest struct {
		SoftwareName *string    `json:"softwareName"`
		PlanName     *string    `json:"planName"`
		Status       *string    `json:"status"`
		TotalSeats   *int       `json:"totalSeats"`
		UsedSeats    *int       `json:"usedSeats"`
		RenewalDate  *time.Time `json:"renewalDate"`
		AnnualCost   *float64   `json:"annualCost"`
	}

	var payload licenseUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.SoftwareName != nil {
		item.SoftwareName = trimSpace(*payload.SoftwareName)
	}
	if payload.PlanName != nil {
		item.PlanName = trimSpace(*payload.PlanName)
	}
	if payload.Status != nil {
		item.Status = trimSpace(*payload.Status)
	}
	if payload.TotalSeats != nil {
		item.TotalSeats = *payload.TotalSeats
	}
	if payload.UsedSeats != nil {
		item.UsedSeats = *payload.UsedSeats
	}
	if payload.RenewalDate != nil {
		item.RenewalDate = *payload.RenewalDate
	}
	if payload.AnnualCost != nil {
		item.AnnualCost = *payload.AnnualCost
	}

	if !requireNonEmpty(c, "softwareName", item.SoftwareName) ||
		!requireNonEmpty(c, "planName", item.PlanName) {
		return
	}
	if item.Status == "" {
		item.Status = "Active"
	}
	if !validateStatus(c, "status", item.Status, []string{"Active", "Expiring Soon", "Expired"}) {
		return
	}
	if !requireIntPositiveOrZero(c, "totalSeats", item.TotalSeats) ||
		!requireIntPositiveOrZero(c, "usedSeats", item.UsedSeats) {
		return
	}
	if item.UsedSeats > item.TotalSeats {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usedSeats cannot exceed totalSeats"})
		return
	}
	if item.RenewalDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "renewalDate is required"})
		return
	}
	if !requirePositiveOrZero(c, "annualCost", item.AnnualCost) {
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteSoftwareLicense(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.SoftwareLicense{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func DeleteDepreciation(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.DepreciationSchedule{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
