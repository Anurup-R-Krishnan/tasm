package handlers

import (
	"net/http"
	"time"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

// GetDashboardStats returns aggregated metrics for the main dashboard
func GetDashboardStats(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var totalAssets int64
	db.Model(&models.Asset{}).Count(&totalAssets)

	var checkedOutAssets int64
	db.Model(&models.Asset{}).Where("status = ?", "Checked Out").Count(&checkedOutAssets)

	var inStockAssets int64
	db.Model(&models.Asset{}).Where("status = ?", "In Stock").Count(&inStockAssets)

	var retiredAssets int64
	db.Model(&models.Asset{}).Where("status = ?", "Retired").Count(&retiredAssets)

	var totalValue float64
	db.Model(&models.Asset{}).Select("COALESCE(SUM(value), 0)").Scan(&totalValue)

	var activeAudits int64
	db.Model(&models.AuditSession{}).Where("status = ?", "Active").Count(&activeAudits)

	var openDiscrepancies int64
	db.Model(&models.AuditDiscrepancy{}).Where("status = ? OR status = ?", "Open", "").Count(&openDiscrepancies)

	var openWorkOrders int64
	db.Model(&models.WorkOrder{}).Where("status IN ?", []string{"Open", "In Progress"}).Count(&openWorkOrders)

	var pendingProcurements int64
	db.Model(&models.ProcurementRequest{}).Where("status = ?", "Pending Approval").Count(&pendingProcurements)

	var expiringWarranties int64
	thirtyDays := time.Now().AddDate(0, 0, 30)
	db.Model(&models.Asset{}).Where("warranty_expiry <= ? AND warranty_expiry > ?", thirtyDays, time.Now()).Count(&expiringWarranties)

	var unreadAlerts int64
	db.Model(&models.SystemAlert{}).Where("is_read = ?", false).Count(&unreadAlerts)

	var lowStockConsumables int64
	db.Model(&models.Consumable{}).Where("current_stock <= reorder_level").Count(&lowStockConsumables)

	var activeReservations int64
	db.Model(&models.Reservation{}).Where("status = ? AND end_time > ?", "Active", time.Now()).Count(&activeReservations)

	var totalUsers int64
	db.Model(&models.SystemUser{}).Where("status = ?", "Active").Count(&totalUsers)

	c.JSON(http.StatusOK, gin.H{
		"totalAssets":         totalAssets,
		"checkedOutAssets":    checkedOutAssets,
		"inStockAssets":       inStockAssets,
		"retiredAssets":       retiredAssets,
		"totalValue":          totalValue,
		"activeAudits":        activeAudits,
		"openDiscrepancies":   openDiscrepancies,
		"openWorkOrders":      openWorkOrders,
		"pendingProcurements": pendingProcurements,
		"expiringWarranties":  expiringWarranties,
		"unreadAlerts":        unreadAlerts,
		"lowStockConsumables": lowStockConsumables,
		"activeReservations":  activeReservations,
		"totalUsers":          totalUsers,
	})
}
