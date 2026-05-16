package handlers

import (
	"net/http"
	"strings"
	"time"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetContracts(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var contracts []models.MaintenanceContract
	query := db.Order("start_date desc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if vendor := strings.TrimSpace(c.Query("vendor")); vendor != "" {
		query = query.Where("vendor_name ILIKE ?", "%"+vendor+"%")
	}
	if err := query.Find(&contracts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load contracts"})
		return
	}

	c.JSON(http.StatusOK, contracts)
}

func GetContractByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var contract models.MaintenanceContract
	if err := db.First(&contract, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}

	c.JSON(http.StatusOK, contract)
}

func CreateContract(c *gin.Context) {
	var contract models.MaintenanceContract
	if !bindJSON(c, &contract) {
		return
	}

	contract.VendorName = trimSpace(contract.VendorName)
	contract.ServiceType = trimSpace(contract.ServiceType)
	contract.Status = trimSpace(contract.Status)

	if !requireNonEmpty(c, "vendorName", contract.VendorName) ||
		!requireNonEmpty(c, "serviceType", contract.ServiceType) {
		return
	}

	if contract.StartDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate is required"})
		return
	}
	if contract.EndDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endDate is required"})
		return
	}
	if contract.EndDate.Before(contract.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endDate cannot be before startDate"})
		return
	}

	if contract.Status == "" {
		contract.Status = "Active"
	}
	if !validateStatus(c, "status", contract.Status, []string{"Active", "Expired"}) {
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&contract).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contract"})
		return
	}

	c.JSON(http.StatusCreated, contract)
}

func GetWorkOrders(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var workOrders []models.WorkOrder
	query := db.Order("target_date asc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if severity := strings.TrimSpace(c.Query("severity")); severity != "" {
		query = query.Where("severity = ?", severity)
	}
	if err := query.Find(&workOrders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load work orders"})
		return
	}

	c.JSON(http.StatusOK, workOrders)
}

func GetWorkOrderByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var workOrder models.WorkOrder
	if err := db.First(&workOrder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Work order not found"})
		return
	}

	c.JSON(http.StatusOK, workOrder)
}

func CreateWorkOrder(c *gin.Context) {
	type workOrderRequest struct {
		WorkOrderID   string  `json:"workOrderId"`
		Title         string  `json:"title"`
		AssetLocation string  `json:"assetLocation"`
		AssetName     string  `json:"assetName"`
		AssetTag      string  `json:"assetTag"`
		Description   string  `json:"description"`
		Issue         string  `json:"issue"`
		Severity      string  `json:"severity"`
		ReportedBy    string  `json:"reportedBy"`
		TargetDate    string  `json:"targetDate"`
		Status        string  `json:"status"`
		Technician    string  `json:"technician"`
		Cost          float64 `json:"cost"`
		EstimatedCost float64 `json:"estimatedCost"`
	}

	var payload workOrderRequest
	if !bindJSON(c, &payload) {
		return
	}

	payload.WorkOrderID = trimSpace(payload.WorkOrderID)
	payload.Title = trimSpace(payload.Title)
	payload.AssetLocation = trimSpace(payload.AssetLocation)
	payload.AssetName = trimSpace(payload.AssetName)
	payload.AssetTag = trimSpace(payload.AssetTag)
	payload.Description = trimSpace(payload.Description)
	payload.Issue = trimSpace(payload.Issue)
	payload.Severity = trimSpace(payload.Severity)
	payload.ReportedBy = trimSpace(payload.ReportedBy)
	payload.Status = trimSpace(payload.Status)
	payload.Technician = trimSpace(payload.Technician)

	if !requireNonEmpty(c, "workOrderId", payload.WorkOrderID) ||
		!requireNonEmpty(c, "title", payload.Title) ||
		!requireNonEmpty(c, "assetLocation", payload.AssetLocation) ||
		!requireNonEmpty(c, "issue", payload.Issue) {
		return
	}

	targetDate, err := parseTime(payload.TargetDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "targetDate must be a valid datetime"})
		return
	}

	if payload.Severity == "" {
		payload.Severity = "Medium"
	}
	if !validateStatus(c, "severity", payload.Severity, []string{"Critical", "High", "Medium", "Low"}) {
		return
	}

	if payload.Status == "" {
		payload.Status = "Open"
	}
	if !validateStatus(c, "status", payload.Status, []string{"Open", "In Progress", "Closed", "On Hold"}) {
		return
	}

	if payload.EstimatedCost == 0 && payload.Cost != 0 {
		payload.EstimatedCost = payload.Cost
	}

	if !requirePositiveOrZero(c, "cost", payload.Cost) ||
		!requirePositiveOrZero(c, "estimatedCost", payload.EstimatedCost) {
		return
	}

	workOrder := models.WorkOrder{
		WorkOrderID:   payload.WorkOrderID,
		Title:         payload.Title,
		AssetLocation: payload.AssetLocation,
		AssetName:     payload.AssetName,
		AssetTag:      payload.AssetTag,
		Description:   payload.Description,
		Issue:         payload.Issue,
		Severity:      payload.Severity,
		ReportedBy:    payload.ReportedBy,
		TargetDate:    targetDate,
		Status:        payload.Status,
		Technician:    payload.Technician,
		Cost:          payload.Cost,
		EstimatedCost: payload.EstimatedCost,
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&workOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create work order"})
		return
	}

	c.JSON(http.StatusCreated, workOrder)
}

func UpdateContract(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.MaintenanceContract
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	var payload models.MaintenanceContract
	if !bindJSON(c, &payload) {
		return
	}

	if payload.VendorName != "" {
		item.VendorName = trimSpace(payload.VendorName)
	}
	if payload.ServiceType != "" {
		item.ServiceType = trimSpace(payload.ServiceType)
	}
	if !payload.StartDate.IsZero() {
		item.StartDate = payload.StartDate
	}
	if !payload.EndDate.IsZero() {
		item.EndDate = payload.EndDate
	}
	if payload.Status != "" {
		item.Status = trimSpace(payload.Status)
	}

	if !requireNonEmpty(c, "vendorName", item.VendorName) ||
		!requireNonEmpty(c, "serviceType", item.ServiceType) {
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

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteContract(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}
	deleteEntity(c, db, &models.MaintenanceContract{}, id)
}

func UpdateWorkOrder(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.WorkOrder
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	if !updateWorkOrderFromPayload(c, &item) {
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func updateWorkOrderFromPayload(c *gin.Context, item *models.WorkOrder) bool {
	var payload struct {
		WorkOrderID   *string    `json:"workOrderId"`
		Title         *string    `json:"title"`
		AssetLocation *string    `json:"assetLocation"`
		AssetName     *string    `json:"assetName"`
		AssetTag      *string    `json:"assetTag"`
		Description   *string    `json:"description"`
		Issue         *string    `json:"issue"`
		Severity      *string    `json:"severity"`
		ReportedBy    *string    `json:"reportedBy"`
		TargetDate    *time.Time `json:"targetDate"`
		Status        *string    `json:"status"`
		Technician    *string    `json:"technician"`
		Cost          *float64   `json:"cost"`
		EstimatedCost *float64   `json:"estimatedCost"`
	}

	if !bindJSON(c, &payload) {
		return false
	}

	applyWOStringField(&item.WorkOrderID, payload.WorkOrderID)
	applyWOStringField(&item.Title, payload.Title)
	applyWOStringField(&item.AssetLocation, payload.AssetLocation)
	applyWOStringField(&item.AssetName, payload.AssetName)
	applyWOStringField(&item.AssetTag, payload.AssetTag)
	applyWOStringField(&item.Description, payload.Description)
	applyWOStringField(&item.Issue, payload.Issue)
	applyWOStringField(&item.Severity, payload.Severity)
	applyWOStringField(&item.ReportedBy, payload.ReportedBy)
	applyWOStringField(&item.Status, payload.Status)
	applyWOStringField(&item.Technician, payload.Technician)

	if payload.TargetDate != nil {
		item.TargetDate = *payload.TargetDate
	}
	if payload.Cost != nil {
		item.Cost = *payload.Cost
	}
	if payload.EstimatedCost != nil {
		item.EstimatedCost = *payload.EstimatedCost
	}

	return validateWorkOrderFields(c, item)
}

func applyWOStringField(target *string, src *string) {
	if src != nil {
		*target = trimSpace(*src)
	}
}

func validateWorkOrderFields(c *gin.Context, item *models.WorkOrder) bool {
	if !requireNonEmpty(c, "workOrderId", item.WorkOrderID) ||
		!requireNonEmpty(c, "title", item.Title) ||
		!requireNonEmpty(c, "assetLocation", item.AssetLocation) ||
		!requireNonEmpty(c, "issue", item.Issue) {
		return false
	}

	if item.TargetDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "targetDate is required"})
		return false
	}

	if item.Severity == "" {
		item.Severity = "Medium"
	}
	if !validateStatus(c, "severity", item.Severity, []string{"Critical", "High", "Medium", "Low"}) {
		return false
	}

	if item.Status == "" {
		item.Status = "Open"
	}
	if !validateStatus(c, "status", item.Status, []string{"Open", "In Progress", "Closed", "On Hold"}) {
		return false
	}

	if item.EstimatedCost == 0 && item.Cost != 0 {
		item.EstimatedCost = item.Cost
	}

	if !requirePositiveOrZero(c, "cost", item.Cost) ||
		!requirePositiveOrZero(c, "estimatedCost", item.EstimatedCost) {
		return false
	}

	return true
}

func DeleteWorkOrder(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}
	deleteEntity(c, db, &models.WorkOrder{}, id)
}
