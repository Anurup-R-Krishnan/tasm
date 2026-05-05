package handlers

import (
	"net/http"
	"strings"
	"time"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
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

type assetRequest struct {
	Name           string  `json:"name"`
	TagID          string  `json:"tagId"`
	Category       string  `json:"category"`
	Location       string  `json:"location"`
	Condition      string  `json:"condition"`
	Status         string  `json:"status"`
	Custodian      string  `json:"custodian"`
	CurrentStock   int     `json:"currentStock"`
	ReorderLevel   int     `json:"reorderLevel"`
	Value          float64 `json:"value"`
	PurchaseDate   string  `json:"purchaseDate"`
	WarrantyStatus string  `json:"warrantyStatus"`
	WarrantyExpiry string  `json:"warrantyExpiry"`
}

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

	if !validateStatus(c, "status", payload.Status, []string{"In Stock", "Checked Out", "Reserved", "Retired"}) {
		return
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

	asset := models.Asset{
		Name:           payload.Name,
		TagID:          payload.TagID,
		Category:       payload.Category,
		Location:       payload.Location,
		Condition:      payload.Condition,
		Status:         payload.Status,
		Custodian:      payload.Custodian,
		CurrentStock:   payload.CurrentStock,
		ReorderLevel:   payload.ReorderLevel,
		Value:          payload.Value,
		PurchaseDate:   purchaseDate,
		WarrantyStatus: payload.WarrantyStatus,
		WarrantyExpiry: warrantyExpiry,
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

	type assetUpdateRequest struct {
		Name           *string  `json:"name"`
		TagID          *string  `json:"tagId"`
		Category       *string  `json:"category"`
		Location       *string  `json:"location"`
		Condition      *string  `json:"condition"`
		Status         *string  `json:"status"`
		Custodian      *string  `json:"custodian"`
		CurrentStock   *int     `json:"currentStock"`
		ReorderLevel   *int     `json:"reorderLevel"`
		Value          *float64 `json:"value"`
		PurchaseDate   *string  `json:"purchaseDate"`
		WarrantyStatus *string  `json:"warrantyStatus"`
		WarrantyExpiry *string  `json:"warrantyExpiry"`
	}

	var payload assetUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.Name != nil {
		asset.Name = trimSpace(*payload.Name)
	}
	if payload.TagID != nil {
		asset.TagID = trimSpace(*payload.TagID)
	}
	if payload.Category != nil {
		asset.Category = trimSpace(*payload.Category)
	}
	if payload.Location != nil {
		asset.Location = trimSpace(*payload.Location)
	}
	if payload.Condition != nil {
		asset.Condition = trimSpace(*payload.Condition)
	}
	if payload.Status != nil {
		asset.Status = trimSpace(*payload.Status)
	}
	if payload.Custodian != nil {
		asset.Custodian = trimSpace(*payload.Custodian)
	}
	if payload.CurrentStock != nil {
		asset.CurrentStock = *payload.CurrentStock
	}
	if payload.ReorderLevel != nil {
		asset.ReorderLevel = *payload.ReorderLevel
	}
	if payload.Value != nil {
		asset.Value = *payload.Value
	}
	if payload.PurchaseDate != nil {
		parsed, err := parseTime(*payload.PurchaseDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "purchaseDate must be a valid datetime"})
			return
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
			return
		}
		asset.WarrantyExpiry = parsed
	}

	if asset.Name == "" || asset.TagID == "" || asset.Category == "" || asset.Location == "" || asset.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, tagId, category, location, and status are required"})
		return
	}

	if !requireIntPositiveOrZero(c, "currentStock", asset.CurrentStock) ||
		!requireIntPositiveOrZero(c, "reorderLevel", asset.ReorderLevel) ||
		!requirePositiveOrZero(c, "value", asset.Value) {
		return
	}

	if !validateStatus(c, "status", asset.Status, []string{"In Stock", "Checked Out", "Reserved", "Retired"}) {
		return
	}

	if !asset.PurchaseDate.IsZero() && asset.PurchaseDate.After(time.Now().AddDate(1, 0, 0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "purchaseDate is not realistic"})
		return
	}

	if asset.WarrantyExpiry.IsZero() {
		asset.WarrantyExpiry = asset.PurchaseDate.AddDate(1, 0, 0)
	}

	if !asset.PurchaseDate.IsZero() && asset.WarrantyExpiry.Before(asset.PurchaseDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warrantyExpiry cannot be before purchaseDate"})
		return
	}

	if payload.TagID != nil {
		var existing models.Asset
		if err := db.Where("tag_id = ? AND id <> ?", asset.TagID, asset.ID).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "asset with this tagId already exists"})
			return
		}
	}

	if err := db.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
		return
	}

	c.JSON(http.StatusOK, asset)
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

	if err := db.Delete(&models.Asset{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete asset"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset deleted successfully"})
}
