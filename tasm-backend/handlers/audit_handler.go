package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetAudits(c *gin.Context) {
	var audits []models.AuditSession
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("id desc").Find(&audits)

	if len(audits) == 0 {
		pastDate := time.Now().AddDate(0, -1, 0)
		audits = []models.AuditSession{
			{Title: "Q3 IT Asset Audit", StartDate: time.Now().AddDate(0, 0, -5), Status: "Active", TotalAssets: 450, ScannedAssets: 320, DiscrepancyCount: 12, AuditorName: "Sarah Jenkins", Progress: 71},
			{Title: "Annual Compliance 2024", StartDate: pastDate.AddDate(0, 0, -15), EndDate: &pastDate, Status: "Completed", TotalAssets: 1200, ScannedAssets: 1200, DiscrepancyCount: 5, AuditorName: "Mike Ross", Progress: 100},
			{Title: "Furniture Inventory Audit", StartDate: time.Now().AddDate(0, 0, -1), Status: "Active", TotalAssets: 800, ScannedAssets: 150, DiscrepancyCount: 2, AuditorName: "Emma Clark", Progress: 18},
		}
		for _, a := range audits {
			database.DB.Create(&a)
		}
		database.DB.Order("id desc").Find(&audits)
	}

	c.JSON(http.StatusOK, audits)
}

func CreateAudit(c *gin.Context) {
	var audit models.AuditSession
	if err := c.ShouldBindJSON(&audit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Create(&audit)
	c.JSON(http.StatusCreated, audit)
}

func GetDiscrepancies(c *gin.Context) {
	var discs []models.AuditDiscrepancy
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Find(&discs)

	if len(discs) == 0 {
		discs = []models.AuditDiscrepancy{
			{AssetTag: "AST-992-LX", IssueType: "Location Mismatch", LastKnownLocation: "C-Wing, Server Rm 2", ScannedLocation: "B-Wing, IT Store", RecommendedAction: "Update to B-Wing"},
			{AssetTag: "MAC-844-PR", IssueType: "Missing", LastKnownLocation: "C-Wing, Desk 104", ScannedLocation: "Not Scanned", RecommendedAction: "Investigate / Mark Lost"},
			{AssetTag: "MON-112-DK", IssueType: "Unregistered", LastKnownLocation: "None", ScannedLocation: "C-Wing, Meeting Rm B", RecommendedAction: "Register New Asset"},
			{AssetTag: "AST-945-LX", IssueType: "Location Mismatch", LastKnownLocation: "C-Wing, Desk 88", ScannedLocation: "C-Wing, Desk 89", RecommendedAction: "Update to Desk 89"},
			{AssetTag: "SRV-004-PR", IssueType: "Missing", LastKnownLocation: "Data Center Alpha", ScannedLocation: "Not Scanned", RecommendedAction: "High Priority Investigate"},
		}
		for _, d := range discs {
			database.DB.Create(&d)
		}
		database.DB.Find(&discs)
	}

	c.JSON(http.StatusOK, discs)
}
