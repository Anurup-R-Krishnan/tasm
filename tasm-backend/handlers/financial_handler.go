package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetLedgers(c *gin.Context) {
	var ledgers []models.LedgerEntry
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("date desc").Find(&ledgers)

	if len(ledgers) == 0 {
		ledgers = []models.LedgerEntry{
			{TransactionID: "TRX-1001", Date: database.DB.NowFunc().AddDate(0, 0, -2), Description: "Server Rack Purchase", Amount: 15000, Type: "Debit", Category: "Hardware"},
			{TransactionID: "TRX-1002", Date: database.DB.NowFunc().AddDate(0, 0, -5), Description: "Software License Renewal", Amount: 3500, Type: "Debit", Category: "Software"},
			{TransactionID: "TRX-1003", Date: database.DB.NowFunc().AddDate(0, 0, -10), Description: "Old Asset Sale", Amount: 2000, Type: "Credit", Category: "Hardware"},
			{TransactionID: "TRX-1004", Date: database.DB.NowFunc().AddDate(0, -1, 0), Description: "Consulting Fees", Amount: 5000, Type: "Debit", Category: "Services"},
			{TransactionID: "TRX-1005", Date: database.DB.NowFunc().AddDate(0, -1, -15), Description: "Cloud Hosting", Amount: 1200, Type: "Debit", Category: "Infrastructure"},
		}
		for _, l := range ledgers {
			database.DB.Create(&l)
		}
		database.DB.Order("date desc").Find(&ledgers)
	}

	c.JSON(http.StatusOK, ledgers)
}

func GetLeases(c *gin.Context) {
	var leases []models.LeaseAgreement
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("start_date desc").Find(&leases)

	if len(leases) == 0 {
		leases = []models.LeaseAgreement{
			{LeaseID: "LSE-2023-01", Vendor: "Dell EMC", StartDate: database.DB.NowFunc().AddDate(-1, 0, 0), EndDate: database.DB.NowFunc().AddDate(2, 0, 0), MonthlyCost: 450.00, Status: "Active"},
			{LeaseID: "LSE-2023-02", Vendor: "Cisco Systems", StartDate: database.DB.NowFunc().AddDate(0, -6, 0), EndDate: database.DB.NowFunc().AddDate(1, 6, 0), MonthlyCost: 1200.00, Status: "Active"},
			{LeaseID: "LSE-2021-01", Vendor: "HP Enterprise", StartDate: database.DB.NowFunc().AddDate(-3, 0, 0), EndDate: database.DB.NowFunc().AddDate(-1, 0, 0), MonthlyCost: 800.00, Status: "Expired"},
		}
		for _, l := range leases {
			database.DB.Create(&l)
		}
		database.DB.Order("start_date desc").Find(&leases)
	}

	c.JSON(http.StatusOK, leases)
}

func GetDepreciations(c *gin.Context) {
	var schedules []models.DepreciationSchedule
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("id desc").Find(&schedules)

	if len(schedules) == 0 {
		schedules = []models.DepreciationSchedule{
			{AssetID: "AST-2023-001", AssetName: "Dell PowerEdge Server", PurchaseValue: 5000.00, CurrentValue: 3500.00, Method: "Straight Line"},
			{AssetID: "AST-2023-002", AssetName: "Cisco Core Switch", PurchaseValue: 12000.00, CurrentValue: 9000.00, Method: "Declining Balance"},
			{AssetID: "AST-2022-045", AssetName: "Office Chair batch (50)", PurchaseValue: 2500.00, CurrentValue: 1500.00, Method: "Straight Line"},
			{AssetID: "AST-2021-112", AssetName: "Conference Room Display", PurchaseValue: 3000.00, CurrentValue: 500.00, Method: "Straight Line"},
		}
		for _, s := range schedules {
			database.DB.Create(&s)
		}
		database.DB.Order("id desc").Find(&schedules)
	}

	c.JSON(http.StatusOK, schedules)
}
