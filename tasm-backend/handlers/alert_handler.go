package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlerts(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var alerts []models.SystemAlert
	query := db.Order("created_at desc")
	if alertType := strings.TrimSpace(c.Query("type")); alertType != "" {
		query = query.Where("type = ?", alertType)
	}
	if c.Query("unread") == "true" {
		query = query.Where("is_read = false")
	}
	if err := query.Find(&alerts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load alerts"})
		return
	}

	c.JSON(http.StatusOK, alerts)
}

func GetAlertByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var alert models.SystemAlert
	if err := db.First(&alert, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	c.JSON(http.StatusOK, alert)
}

func UpdateAlert(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.SystemAlert
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	type alertUpdateRequest struct {
		Title   *string `json:"title"`
		Message *string `json:"message"`
		Type    *string `json:"type"`
		Source  *string `json:"source"`
		IsRead  *bool   `json:"isRead"`
	}

	var payload alertUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.Title != nil {
		item.Title = trimSpace(*payload.Title)
	}
	if payload.Message != nil {
		item.Message = trimSpace(*payload.Message)
	}
	if payload.Type != nil {
		item.Type = trimSpace(*payload.Type)
	}
	if payload.Source != nil {
		item.Source = trimSpace(*payload.Source)
	}
	if payload.IsRead != nil {
		item.IsRead = *payload.IsRead
	}

	if !requireNonEmpty(c, "title", item.Title) ||
		!requireNonEmpty(c, "message", item.Message) ||
		!requireNonEmpty(c, "type", item.Type) {
		return
	}
	if !validateStatus(c, "type", item.Type, []string{"Critical", "Warning", "Info"}) {
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteAlert(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Delete(&models.SystemAlert{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

// MarkAllAlertsRead marks all unread alerts as read
func MarkAllAlertsRead(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}
	if err := db.Model(&models.SystemAlert{}).Where("is_read = ?", false).Update("is_read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark alerts as read"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All alerts marked as read"})
}

// GenerateAlerts scans the database for conditions that warrant alerts and creates them
func GenerateAlerts(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}
	count := generateSystemAlerts(db)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d new alerts generated", count)})
}

// generateSystemAlerts is the core logic for auto-generating alerts (reusable)
func generateSystemAlerts(db *gorm.DB) int {
	count := 0
	now := time.Now()
	soon := now.AddDate(0, 0, 30)

	// 1. Expiring warranties (within 30 days)
	var expiringWarrantyAssets []models.Asset
	db.Where("warranty_expiry > ? AND warranty_expiry <= ?", now, soon).Find(&expiringWarrantyAssets)
	for _, asset := range expiringWarrantyAssets {
		title := "Warranty Expiring: " + asset.Name
		var existing models.SystemAlert
		if db.Where("title = ? AND is_read = ?", title, false).First(&existing).Error != nil {
			db.Create(&models.SystemAlert{
				Title:   title,
				Message: fmt.Sprintf("Warranty for %s (TAG: %s) expires on %s", asset.Name, asset.TagID, asset.WarrantyExpiry.Format("02 Jan 2006")),
				Type:    "Warning",
				Source:  "Warranty Tracker",
				IsRead:  false,
			})
			count++
		}
	}

	// 2. Low stock consumables
	var lowStockConsumables []models.Consumable
	db.Where("current_stock <= reorder_level AND reorder_level > 0").Find(&lowStockConsumables)
	for _, c := range lowStockConsumables {
		title := "Low Stock: " + c.Name
		var existing models.SystemAlert
		if db.Where("title = ? AND is_read = ?", title, false).First(&existing).Error != nil {
			db.Create(&models.SystemAlert{
				Title:   title,
				Message: fmt.Sprintf("%s stock is at %d units (reorder level: %d) at %s", c.Name, c.CurrentStock, c.ReorderLevel, c.Location),
				Type:    "Warning",
				Source:  "Inventory Tracker",
				IsRead:  false,
			})
			count++
		}
	}

	// 3. Expiring software licenses (within 30 days)
	var expLicenses []models.SoftwareLicense
	db.Where("renewal_date > ? AND renewal_date <= ?", now, soon).Find(&expLicenses)
	for _, lic := range expLicenses {
		title := "License Renewal: " + lic.SoftwareName
		var existing models.SystemAlert
		if db.Where("title = ? AND is_read = ?", title, false).First(&existing).Error != nil {
			alertType := "Warning"
			if lic.RenewalDate.Before(now.AddDate(0, 0, 7)) {
				alertType = "Critical"
			}
			db.Create(&models.SystemAlert{
				Title:   title,
				Message: fmt.Sprintf("%s (%s) license renewal due on %s. Annual cost: $%.2f", lic.SoftwareName, lic.PlanName, lic.RenewalDate.Format("02 Jan 2006"), lic.AnnualCost),
				Type:    alertType,
				Source:  "License Tracker",
				IsRead:  false,
			})
			count++
		}
	}

	// 4. Overdue work orders (target date passed and still open)
	var overdueOrders []models.WorkOrder
	db.Where("target_date < ? AND status IN ?", now, []string{"Open", "In Progress"}).Find(&overdueOrders)
	for _, wo := range overdueOrders {
		title := "Overdue Work Order: " + wo.WorkOrderID
		var existing models.SystemAlert
		if db.Where("title = ? AND is_read = ?", title, false).First(&existing).Error != nil {
			alertType := "Warning"
			if wo.Severity == "Critical" || wo.Severity == "High" {
				alertType = "Critical"
			}
			db.Create(&models.SystemAlert{
				Title:   title,
				Message: fmt.Sprintf("Work order %s (%s) is overdue. Assigned to: %s", wo.WorkOrderID, wo.Title, wo.Technician),
				Type:    alertType,
				Source:  "Maintenance Tracker",
				IsRead:  false,
			})
			count++
		}
	}

	return count
}
