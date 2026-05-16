package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAssets(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var assets []models.Asset
	query := db.Order("created_at desc")

	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}

	if category := strings.TrimSpace(c.Query("category")); category != "" {
		query = query.Where("category = ?", category)
	}

	if location := strings.TrimSpace(c.Query("location")); location != "" {
		query = query.Where("location = ?", location)
	}

	if custodian := strings.TrimSpace(c.Query("custodian")); custodian != "" {
		query = query.Where("custodian = ?", custodian)
	}

	if search := strings.TrimSpace(c.Query("search")); search != "" {
		like := "%" + search + "%"
		query = query.Where(
			"name ILIKE ? OR tag_id ILIKE ? OR location ILIKE ? OR custodian ILIKE ?",
			like,
			like,
			like,
			like,
		)
	}

	if err := query.Find(&assets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load assets"})
		return
	}

	c.JSON(http.StatusOK, assets)
}

// CheckoutAsset handles checking out an asset to a user
func CheckoutAsset(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	if asset.Status != "In Stock" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Asset is not available for checkout"})
		return
	}

	type checkoutPayload struct {
		UserID    uint   `json:"userId"`
		DueDate   string `json:"dueDate"` // ISO date string
		Notes     string `json:"notes,omitempty"`
		Custodian string `json:"custodian"` // optional override
	}
	var payload checkoutPayload
	if !bindJSON(c, &payload) {
		return
	}

	if payload.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	var user models.SystemUser
	if err := db.First(&user, payload.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	if strings.TrimSpace(user.Status) != "Active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is not active"})
		return
	}

	var custodian string
	if strings.TrimSpace(payload.Custodian) != "" {
		custodian = strings.TrimSpace(payload.Custodian)
	} else {
		custodian = user.Name
	}

	var dueDate *time.Time
	if strings.TrimSpace(payload.DueDate) != "" {
		parsed, err := parseTime(payload.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "dueDate must be a valid datetime"})
			return
		}
		dueDate = &parsed
		if dueDate.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "dueDate cannot be in the past"})
			return
		}
	}

	// Update asset
	asset.Status = "Checked Out"
	asset.Custodian = custodian
	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
		return
	}

	// Log event
	logAssetEvent(db, asset.ID, user.ID, user.Name, "CHECKOUT",
		"Asset checked out",
		"In Stock", asset.Status,
		"", asset.Custodian,
		dueDate, payload.Notes)

	c.JSON(http.StatusOK, asset)
}

// CheckinAsset handles checking in an asset
func CheckinAsset(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	if asset.Status != "Checked Out" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Asset is not currently checked out"})
		return
	}

	type checkinPayload struct {
		Notes     string `json:"notes,omitempty"`
		Custodian string `json:"custodian,omitempty"` // optional, for verification
	}
	var payload checkinPayload
	if !bindJSON(c, &payload) {
		return
	}

	previousCustodian := asset.Custodian
	if strings.TrimSpace(payload.Custodian) != "" && strings.TrimSpace(payload.Custodian) != previousCustodian {
		c.JSON(http.StatusBadRequest, gin.H{"error": "custodian does not match current holder"})
		return
	}

	// Update asset
	asset.Status = "In Stock"
	asset.Custodian = ""
	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
		return
	}

	// Log event - we need actor from context (the user performing checkin)
	userID, exists := c.Get("userID")
	var actorID uint
	var actorName string
	if exists {
		actorID = userID.(uint)
		var actor models.SystemUser
		if err := db.First(&actor, actorID).Error; err == nil {
			actorName = actor.Name
		}
	}
	if actorName == "" {
		actorName = "System"
	}

	logAssetEvent(db, asset.ID, actorID, actorName, "CHECKIN",
		"Asset checked in",
		"Checked Out", asset.Status,
		previousCustodian, "",
		nil, payload.Notes)

	c.JSON(http.StatusOK, asset)
}

// TransferAsset handles transferring an asset to a new location or custodian
func TransferAsset(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	type transferPayload struct {
		NewLocation  string `json:"newLocation"`
		NewCustodian string `json:"newCustodian"`
		Notes        string `json:"notes,omitempty"`
	}
	var payload transferPayload
	if !bindJSON(c, &payload) {
		return
	}

	previousLocation := asset.Location
	previousCustodian := asset.Custodian

	if payload.NewLocation != "" {
		asset.Location = payload.NewLocation
	}
	if payload.NewCustodian != "" {
		asset.Custodian = payload.NewCustodian
	}

	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset during transfer"})
		return
	}

	// Log event
	userID, exists := c.Get("userID")
	var actorID uint
	var actorName string
	if exists {
		actorID = userID.(uint)
		var actor models.SystemUser
		if err := db.First(&actor, actorID).Error; err == nil {
			actorName = actor.Name
		}
	}
	if actorName == "" {
		actorName = "System"
	}

	description := "Asset transferred"
	if payload.NewLocation != "" && payload.NewLocation != previousLocation {
		description += " to " + payload.NewLocation
	}
	if payload.NewCustodian != "" && payload.NewCustodian != previousCustodian {
		description += " to " + payload.NewCustodian
	}

	logAssetEvent(db, asset.ID, actorID, actorName, "TRANSFER",
		description,
		asset.Status, asset.Status,
		previousCustodian, asset.Custodian,
		nil, payload.Notes)

	c.JSON(http.StatusOK, asset)
}

// GetAssetHistory returns the event history for an asset
func GetAssetHistory(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var events []models.AssetEvent
	if err := db.Where("asset_id = ?", id).Order("created_at desc").Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch asset history"})
		return
	}

	c.JSON(http.StatusOK, events)
}

// GetAssetByTag returns the asset record matching the given tag ID
func GetAssetByTag(c *gin.Context) {
	tagID := c.Param("tagId")
	if tagID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tagId is required"})
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.Where("tag_id = ?", tagID).First(&asset).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "asset not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch asset"})
		return
	}

	c.JSON(http.StatusOK, asset)
}

// logAssetEvent creates an asset event record
func logAssetEvent(db *gorm.DB, assetID uint, actorID uint, actorName string, eventType string, description string,
	previousStatus string, newStatus string,
	previousCustodian string, newCustodian string,
	dueDate *time.Time, notes string) {
	logAssetEventFull(db, assetID, actorID, actorName, eventType, description,
		previousStatus, newStatus, previousCustodian, newCustodian, "", "", dueDate, notes)
}

func logAssetEventFull(db *gorm.DB, assetID uint, actorID uint, actorName string, eventType string, description string,
	previousStatus string, newStatus string,
	previousCustodian string, newCustodian string,
	previousLocation string, newLocation string,
	dueDate *time.Time, notes string) {
	event := models.AssetEvent{
		AssetID:           assetID,
		EventType:         eventType,
		ActorID:           actorID,
		ActorName:         actorName,
		Description:       description,
		PreviousStatus:    previousStatus,
		NewStatus:         newStatus,
		PreviousCustodian: previousCustodian,
		NewCustodian:      newCustodian,
		PreviousLocation:  previousLocation,
		NewLocation:       newLocation,
		DueDate:           dueDate,
		Notes:             notes,
	}
	db.Create(&event)
}

type assetRequest struct {
	Name               string  `json:"name"`
	TagID              string  `json:"tagId"`
	SerialNumber       string  `json:"serialNumber"`
	Category           string  `json:"category"`
	Location           string  `json:"location"`
	Condition          string  `json:"condition"`
	Status             string  `json:"status"`
	Custodian          string  `json:"custodian"`
	CurrentStock       int     `json:"currentStock"`
	ReorderLevel       int     `json:"reorderLevel"`
	Value              float64 `json:"value"`
	PurchaseDate       string  `json:"purchaseDate"`
	PurchasePrice      float64 `json:"purchasePrice"`
	WarrantyStatus     string  `json:"warrantyStatus"`
	WarrantyExpiry     string  `json:"warrantyExpiry"`
	WarrantyStartDate  string  `json:"warrantyStartDate"`
	WarrantyEndDate    string  `json:"warrantyEndDate"`
	WarrantyProvider   string  `json:"warrantyProvider"`
	WarrantyTerms      string  `json:"warrantyTerms"`
	UsefulLifeYears    int     `json:"usefulLifeYears"`
	DepreciationMethod string  `json:"depreciationMethod"`
	ResidualValue      float64 `json:"residualValue"`
	ReplacementCost    float64 `json:"replacementCost"`
	LifecycleStatus    string  `json:"lifecycleStatus"`
	Department         string  `json:"department"`
}

var validLifecycleStatuses = []string{"Procurement", "Deployed", "Under Maintenance", "End of Life", "Disposed"}
var validAssetStatuses = []string{"In Stock", "Checked Out", "Reserved", "Deployed", "Under Maintenance", "End of Life", "Disposed", "Retired"}

func CreateAsset(c *gin.Context) {
	var payload assetRequest
	if !bindJSON(c, &payload) {
		return
	}

	payload.Name = trimSpace(payload.Name)
	payload.TagID = trimSpace(payload.TagID)
	payload.Category = trimSpace(payload.Category)
	payload.Location = trimSpace(payload.Location)
	payload.Condition = trimSpace(payload.Condition)
	payload.Status = trimSpace(payload.Status)
	payload.Custodian = trimSpace(payload.Custodian)
	payload.WarrantyStatus = trimSpace(payload.WarrantyStatus)
	payload.DepreciationMethod = trimSpace(payload.DepreciationMethod)
	payload.LifecycleStatus = trimSpace(payload.LifecycleStatus)

	if !requireNonEmpty(c, "name", payload.Name) ||
		!requireNonEmpty(c, "tagId", payload.TagID) ||
		!requireNonEmpty(c, "category", payload.Category) ||
		!requireNonEmpty(c, "location", payload.Location) ||
		!requireNonEmpty(c, "status", payload.Status) {
		return
	}

	if !requireIntPositiveOrZero(c, "currentStock", payload.CurrentStock) ||
		!requireIntPositiveOrZero(c, "reorderLevel", payload.ReorderLevel) ||
		!requirePositiveOrZero(c, "value", payload.Value) {
		return
	}

	if !validateStatus(c, "status", payload.Status, validAssetStatuses) {
		return
	}

	// Default depreciation method
	if payload.DepreciationMethod == "" {
		payload.DepreciationMethod = "Straight Line"
	}
	if payload.DepreciationMethod != "Straight Line" && payload.DepreciationMethod != "Declining Balance" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "depreciationMethod must be Straight Line or Declining Balance"})
		return
	}

	// Default lifecycle status
	if payload.LifecycleStatus == "" {
		payload.LifecycleStatus = "Deployed"
	}

	// Purchase price defaults to value if not set
	if payload.PurchasePrice <= 0 {
		payload.PurchasePrice = payload.Value
	}

	var purchaseDate time.Time
	if payload.PurchaseDate != "" {
		parsed, err := parseTime(payload.PurchaseDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "purchaseDate must be a valid datetime"})
			return
		}
		purchaseDate = parsed
	} else {
		purchaseDate = time.Now()
	}

	if purchaseDate.After(time.Now().AddDate(1, 0, 0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "purchaseDate is not realistic"})
		return
	}

	var warrantyExpiry time.Time
	if payload.WarrantyExpiry != "" {
		parsed, err := parseTime(payload.WarrantyExpiry)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "warrantyExpiry must be a valid datetime"})
			return
		}
		warrantyExpiry = parsed
	} else {
		warrantyExpiry = purchaseDate.AddDate(1, 0, 0)
	}

	if warrantyExpiry.Before(purchaseDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warrantyExpiry cannot be before purchaseDate"})
		return
	}

	// Parse optional warranty dates
	warrantyStart := purchaseDate
	if payload.WarrantyStartDate != "" {
		if parsed, err := parseTime(payload.WarrantyStartDate); err == nil {
			warrantyStart = parsed
		}
	}
	warrantyEnd := warrantyExpiry
	if payload.WarrantyEndDate != "" {
		if parsed, err := parseTime(payload.WarrantyEndDate); err == nil {
			warrantyEnd = parsed
		}
	}

	asset := models.Asset{
		Name:               payload.Name,
		TagID:              payload.TagID,
		SerialNumber:       payload.SerialNumber,
		Department:         payload.Department,
		Category:           payload.Category,
		Location:           payload.Location,
		Condition:          payload.Condition,
		Status:             payload.Status,
		Custodian:          payload.Custodian,
		CurrentStock:       payload.CurrentStock,
		ReorderLevel:       payload.ReorderLevel,
		Value:              payload.Value,
		PurchaseDate:       purchaseDate,
		PurchasePrice:      payload.PurchasePrice,
		WarrantyStatus:     payload.WarrantyStatus,
		WarrantyExpiry:     warrantyExpiry,
		WarrantyStartDate:  warrantyStart,
		WarrantyEndDate:    warrantyEnd,
		WarrantyProvider:   payload.WarrantyProvider,
		WarrantyTerms:      payload.WarrantyTerms,
		UsefulLifeYears:    payload.UsefulLifeYears,
		DepreciationMethod: payload.DepreciationMethod,
		ResidualValue:      payload.ResidualValue,
		ReplacementCost:    payload.ReplacementCost,
		LifecycleStatus:    payload.LifecycleStatus,
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	var existing models.Asset
	if err := db.Where("tag_id = ?", asset.TagID).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "asset with this tagId already exists"})
		return
	}

	if err := db.Create(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create asset"})
		return
	}

	// Audit log
	actorID, actorName := getActorFromContext(c, db)
	logAssetEvent(db, asset.ID, actorID, actorName, "CREATED",
		"Asset created: "+asset.Name,
		"", asset.Status, "", asset.Custodian, nil, "")

	c.JSON(http.StatusCreated, asset)
}

func GetAssetByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	c.JSON(http.StatusOK, asset)
}

func UpdateAsset(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}
	previousStatus := asset.Status

	if !updateAssetFromPayload(c, db, &asset) {
		return
	}

	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
		return
	}

	// Audit log
	actorID, actorName := getActorFromContext(c, db)
	logAssetEvent(db, asset.ID, actorID, actorName, "UPDATED",
		"Asset updated",
		previousStatus, asset.Status, "", "", nil, "")

	c.JSON(http.StatusOK, asset)
}

func updateAssetFromPayload(c *gin.Context, db *gorm.DB, asset *models.Asset) bool {
	var payload struct {
		Name               *string  `json:"name"`
		TagID              *string  `json:"tagId"`
		SerialNumber       *string  `json:"serialNumber"`
		Category           *string  `json:"category"`
		Location           *string  `json:"location"`
		Condition          *string  `json:"condition"`
		Status             *string  `json:"status"`
		Custodian          *string  `json:"custodian"`
		Department         *string  `json:"department"`
		CurrentStock       *int     `json:"currentStock"`
		ReorderLevel       *int     `json:"reorderLevel"`
		Value              *float64 `json:"value"`
		PurchaseDate       *string  `json:"purchaseDate"`
		PurchasePrice      *float64 `json:"purchasePrice"`
		WarrantyStatus     *string  `json:"warrantyStatus"`
		WarrantyExpiry     *string  `json:"warrantyExpiry"`
		WarrantyStartDate  *string  `json:"warrantyStartDate"`
		WarrantyEndDate    *string  `json:"warrantyEndDate"`
		WarrantyProvider   *string  `json:"warrantyProvider"`
		WarrantyTerms      *string  `json:"warrantyTerms"`
		UsefulLifeYears    *int     `json:"usefulLifeYears"`
		DepreciationMethod *string  `json:"depreciationMethod"`
		ResidualValue      *float64 `json:"residualValue"`
		ReplacementCost    *float64 `json:"replacementCost"`
		LifecycleStatus    *string  `json:"lifecycleStatus"`
	}

	if !bindJSON(c, &payload) {
		return false
	}

	applyAssetStringField(&asset.Name, payload.Name)
	applyAssetStringField(&asset.TagID, payload.TagID)
	applyAssetStringField(&asset.SerialNumber, payload.SerialNumber)
	applyAssetStringField(&asset.Category, payload.Category)
	applyAssetStringField(&asset.Location, payload.Location)
	applyAssetStringField(&asset.Condition, payload.Condition)
	applyAssetStringField(&asset.Status, payload.Status)
	applyAssetStringField(&asset.Custodian, payload.Custodian)
	applyAssetStringField(&asset.Department, payload.Department)
	applyAssetIntField(&asset.CurrentStock, payload.CurrentStock)
	applyAssetIntField(&asset.ReorderLevel, payload.ReorderLevel)
	applyAssetFloatField(&asset.Value, payload.Value)
	applyAssetFloatField(&asset.PurchasePrice, payload.PurchasePrice)
	applyAssetStringField(&asset.WarrantyProvider, payload.WarrantyProvider)
	applyAssetStringField(&asset.WarrantyTerms, payload.WarrantyTerms)
	applyAssetIntField(&asset.UsefulLifeYears, payload.UsefulLifeYears)
	applyAssetStringField(&asset.DepreciationMethod, payload.DepreciationMethod)
	applyAssetFloatField(&asset.ResidualValue, payload.ResidualValue)
	applyAssetFloatField(&asset.ReplacementCost, payload.ReplacementCost)
	applyAssetStringField(&asset.LifecycleStatus, payload.LifecycleStatus)

	if payload.PurchaseDate != nil {
		parsed, err := parseTime(*payload.PurchaseDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "purchaseDate must be a valid datetime"})
			return false
		}
		asset.PurchaseDate = parsed
	}

	if payload.WarrantyStatus != nil {
		asset.WarrantyStatus = trimSpace(*payload.WarrantyStatus)
	}

	if payload.WarrantyExpiry != nil {
		parsed, err := parseTime(*payload.WarrantyExpiry)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "warrantyExpiry must be a valid datetime"})
			return false
		}
		asset.WarrantyExpiry = parsed
	}
	if payload.WarrantyStartDate != nil {
		if parsed, err := parseTime(*payload.WarrantyStartDate); err == nil {
			asset.WarrantyStartDate = parsed
		}
	}
	if payload.WarrantyEndDate != nil {
		if parsed, err := parseTime(*payload.WarrantyEndDate); err == nil {
			asset.WarrantyEndDate = parsed
		}
	}

	return validateAssetFields(c, asset, db, payload.TagID)
}

func applyAssetStringField(target *string, src *string) {
	if src != nil {
		*target = trimSpace(*src)
	}
}

func applyAssetIntField(target *int, src *int) {
	if src != nil {
		*target = *src
	}
}

func applyAssetFloatField(target *float64, src *float64) {
	if src != nil {
		*target = *src
	}
}

func validateAssetFields(c *gin.Context, asset *models.Asset, db *gorm.DB, tagIDChanged *string) bool {
	if asset.Name == "" || asset.TagID == "" || asset.Category == "" || asset.Location == "" || asset.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, tagId, category, location, and status are required"})
		return false
	}

	if !requireIntPositiveOrZero(c, "currentStock", asset.CurrentStock) ||
		!requireIntPositiveOrZero(c, "reorderLevel", asset.ReorderLevel) ||
		!requirePositiveOrZero(c, "value", asset.Value) {
		return false
	}

	if !validateStatus(c, "status", asset.Status, validAssetStatuses) {
		return false
	}

	if !asset.PurchaseDate.IsZero() && asset.PurchaseDate.After(time.Now().AddDate(1, 0, 0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "purchaseDate is not realistic"})
		return false
	}

	if asset.WarrantyExpiry.IsZero() {
		asset.WarrantyExpiry = asset.PurchaseDate.AddDate(1, 0, 0)
	}

	if !asset.PurchaseDate.IsZero() && asset.WarrantyExpiry.Before(asset.PurchaseDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warrantyExpiry cannot be before purchaseDate"})
		return false
	}

	if tagIDChanged != nil {
		var existing models.Asset
		if err := db.Where("tag_id = ? AND id <> ?", asset.TagID, asset.ID).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "asset with this tagId already exists"})
			return false
		}
	}

	return true
}

func DeleteAsset(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	// Audit log before delete
	actorID, actorName := getActorFromContext(c, db)
	logAssetEvent(db, id, actorID, actorName, "DELETED", "Asset soft-deleted", "", "", "", "", nil, "")

	deleteEntity(c, db, &models.Asset{}, id)
}

// ---- Retirement & Disposal ----

func RetireAsset(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	type retirePayload struct {
		Reason string `json:"reason"`
		Notes  string `json:"notes"`
	}
	var payload retirePayload
	if !bindJSON(c, &payload) {
		return
	}

	previousStatus := asset.Status
	asset.Status = "End of Life"
	asset.LifecycleStatus = "End of Life"

	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retire asset"})
		return
	}

	actorID, actorName := getActorFromContext(c, db)
	logAssetEvent(db, asset.ID, actorID, actorName, "RETIRED",
		"Asset retired: "+payload.Reason,
		previousStatus, "End of Life", "", "", nil, payload.Notes)

	c.JSON(http.StatusOK, asset)
}

func DisposeAsset(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	type disposePayload struct {
		DisposalMethod          string  `json:"disposalMethod"`
		ResidualValue           float64 `json:"residualValue"`
		EnvironmentalCompliance bool    `json:"environmentalCompliance"`
		DataWipingConfirmed     bool    `json:"dataWipingConfirmed"`
		Notes                   string  `json:"notes"`
	}
	var payload disposePayload
	if !bindJSON(c, &payload) {
		return
	}

	payload.DisposalMethod = trimSpace(payload.DisposalMethod)
	if !validateStatus(c, "disposalMethod", payload.DisposalMethod, []string{"Sell", "Scrap", "Recycle", "Donate"}) {
		return
	}

	previousStatus := asset.Status
	now := time.Now()
	asset.Status = "Disposed"
	asset.LifecycleStatus = "Disposed"
	asset.DisposalMethod = payload.DisposalMethod
	asset.DisposalDate = &now
	asset.DisposalNotes = payload.Notes
	asset.EnvironmentalCompliance = payload.EnvironmentalCompliance
	asset.DataWipingConfirmed = payload.DataWipingConfirmed

	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dispose asset"})
		return
	}

	actorID, actorName := getActorFromContext(c, db)

	// Create disposal record
	record := models.AssetDisposalRecord{
		AssetID:                 asset.ID,
		DisposalMethod:          payload.DisposalMethod,
		ResidualValue:           payload.ResidualValue,
		DecommissionDate:        now,
		ValuationDate:           now,
		EnvironmentalCompliance: payload.EnvironmentalCompliance,
		DataWipingConfirmed:     payload.DataWipingConfirmed,
		ComplianceVerified:      payload.EnvironmentalCompliance,
		Notes:                   payload.Notes,
		ActorID:                 actorID,
		ActorName:               actorName,
	}
	db.Create(&record)

	logAssetEvent(db, asset.ID, actorID, actorName, "DISPOSED",
		"Asset disposed via "+payload.DisposalMethod,
		previousStatus, "Disposed", "", "", nil, payload.Notes)

	c.JSON(http.StatusOK, asset)
}

// ---- Depreciation Schedule (Computed) ----

type DepreciationYearRow struct {
	Year                    int     `json:"year"`
	OpeningValue            float64 `json:"openingValue"`
	DepreciationAmount      float64 `json:"depreciationAmount"`
	AccumulatedDepreciation float64 `json:"accumulatedDepreciation"`
	ClosingBookValue        float64 `json:"closingBookValue"`
}

func GetAssetDepreciationSchedule(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	purchasePrice := asset.PurchasePrice
	if purchasePrice <= 0 {
		purchasePrice = asset.Value
	}
	usefulLife := asset.UsefulLifeYears
	if usefulLife <= 0 {
		usefulLife = 5 // default
	}
	residual := asset.ResidualValue
	method := asset.DepreciationMethod
	if method == "" {
		method = "Straight Line"
	}

	var rows []DepreciationYearRow
	accumulated := 0.0
	bookValue := purchasePrice

	for year := 1; year <= usefulLife; year++ {
		var depr float64
		if method == "Straight Line" {
			depr = (purchasePrice - residual) / float64(usefulLife)
		} else { // Declining Balance
			rate := 2.0 / float64(usefulLife)
			depr = bookValue * rate
			if bookValue-depr < residual {
				depr = bookValue - residual
			}
		}
		if depr < 0 {
			depr = 0
		}

		accumulated += depr
		closing := bookValue - depr

		rows = append(rows, DepreciationYearRow{
			Year:                    year,
			OpeningValue:            bookValue,
			DepreciationAmount:      depr,
			AccumulatedDepreciation: accumulated,
			ClosingBookValue:        closing,
		})
		bookValue = closing
	}

	c.JSON(http.StatusOK, rows)
}

// ---- Transfer History ----

func GetAssetTransfers(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var transfers []models.AssetTransfer
	if err := db.Where("asset_id = ?", id).Order("created_at desc").Find(&transfers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transfers"})
		return
	}

	c.JSON(http.StatusOK, transfers)
}

func RevertTransfer(c *gin.Context) {
	assetID, ok := parseIDParam(c)
	if !ok {
		return
	}
	transferIDStr := c.Param("transferId")
	if transferIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "transferId is required"})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	var transfer models.AssetTransfer
	if err := db.First(&transfer, transferIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transfer not found"})
		return
	}

	if transfer.AssetID != assetID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transfer does not belong to this asset"})
		return
	}
	if transfer.Reverted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transfer already reverted"})
		return
	}

	// Revert asset location/custodian
	var asset models.Asset
	if err := db.First(&asset, assetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	asset.Location = transfer.FromLocation
	asset.Custodian = transfer.FromEntity
	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revert asset"})
		return
	}

	now := time.Now()
	actorID, actorName := getActorFromContext(c, db)
	transfer.Reverted = true
	transfer.RevertedAt = &now
	transfer.RevertedBy = actorName
	db.Save(&transfer)

	logAssetEventFull(db, asset.ID, actorID, actorName, "TRANSFER_REVERTED",
		"Transfer reverted",
		"", "", transfer.ToEntity, transfer.FromEntity,
		transfer.ToLocation, transfer.FromLocation, nil, "")

	c.JSON(http.StatusOK, transfer)
}

// ---- Receipt Upload ----

func UploadReceipt(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	file, err := c.FormFile("receipt")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receipt file is required"})
		return
	}

	// Save to uploads directory
	filename := strings.ReplaceAll(asset.TagID, "/", "_") + "_receipt_" + file.Filename
	savePath := "handlers/uploads/" + filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save receipt"})
		return
	}

	asset.ReceiptFilePath = savePath
	db.Save(&asset)

	c.JSON(http.StatusOK, gin.H{"path": savePath})
}

func GetReceipt(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	if asset.ReceiptFilePath == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt uploaded"})
		return
	}

	c.File(asset.ReceiptFilePath)
}
