package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetSoftwareLicenses(c *gin.Context) {
	var licenses []models.SoftwareLicense
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}
	
	database.DB.Order("id asc").Find(&licenses)

	if len(licenses) == 0 {
		licenses = []models.SoftwareLicense{
			{SoftwareName: "Adobe Creative Cloud", PlanName: "Enterprise All Apps", Status: "Active", TotalSeats: 100, UsedSeats: 85, RenewalDate: database.DB.NowFunc().AddDate(0, 6, 0), AnnualCost: 84000},
			{SoftwareName: "Microsoft 365", PlanName: "E5 License", Status: "Active", TotalSeats: 500, UsedSeats: 475, RenewalDate: database.DB.NowFunc().AddDate(0, 10, 0), AnnualCost: 228000},
			{SoftwareName: "Figma", PlanName: "Organization Plan", Status: "Expiring Soon", TotalSeats: 50, UsedSeats: 50, RenewalDate: database.DB.NowFunc().AddDate(0, 0, 15), AnnualCost: 27000},
		}
		for _, l := range licenses {
			database.DB.Create(&l)
		}
		database.DB.Order("id asc").Find(&licenses)
	}

	c.JSON(http.StatusOK, licenses)
}
