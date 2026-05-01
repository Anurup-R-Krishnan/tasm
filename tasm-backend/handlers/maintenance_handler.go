package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tasm-backend/models"
)

func GetContracts(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var contracts []models.MaintenanceContract
	if err := db.Order("start_date desc").Find(&contracts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load contracts"})
		return
	}

	if len(contracts) == 0 {
		contracts = []models.MaintenanceContract{
			{
				VendorName:  "TechServe Solutions",
				ServiceType: "IT Equipment Servicing",
				StartDate:   db.NowFunc().AddDate(0, -6, 0),
				EndDate:     db.NowFunc().AddDate(0, 6, 0),
				Status:      "Active",
			},
			{
				VendorName:  "Cooling Systems Inc.",
				ServiceType: "HVAC Maintenance",
				StartDate:   db.NowFunc().AddDate(-1, 0, 0),
				EndDate:     db.NowFunc().AddDate(0, -1, 0),
				Status:      "Expired",
			},
		}
		for _, contract := range contracts {
			db.Create(&contract)
		}
		db.Order("start_date desc").Find(&contracts)
	}

	c.JSON(http.StatusOK, contracts)
}

func CreateContract(c *gin.Context) {
	var contract models.MaintenanceContract
	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	if err := db.Order("target_date asc").Find(&workOrders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load work orders"})
		return
	}

	if len(workOrders) == 0 {
		workOrders = []models.WorkOrder{
			{
				WorkOrderID:   "WO-2023-001",
				Title:         "Fix AC in Server Room",
				AssetLocation: "Server Room A",
				Issue:         "AC unit is leaking water",
				Severity:      "High",
				TargetDate:    db.NowFunc().AddDate(0, 0, 2),
				Status:        "Open",
				Technician:    "",
				Cost:          0,
			},
			{
				WorkOrderID:   "WO-2023-002",
				Title:         "Replace UPS Batteries",
				AssetLocation: "Data Center Alpha",
				Issue:         "UPS batteries are degraded",
				Severity:      "Medium",
				TargetDate:    db.NowFunc().AddDate(0, 1, 0),
				Status:        "In Progress",
				Technician:    "Ramesh K.",
				Cost:          0,
			},
			{
				WorkOrderID:   "WO-2023-003",
				Title:         "Update Router Firmware",
				AssetLocation: "Network Closet",
				Issue:         "Security patch required",
				Severity:      "Low",
				TargetDate:    db.NowFunc().AddDate(0, 0, -7),
				Status:        "Closed",
				Technician:    "Suresh M.",
				Cost:          3200.00,
			},
		}
		for _, wo := range workOrders {
			db.Create(&wo)
		}
		db.Order("target_date asc").Find(&workOrders)
	}

	c.JSON(http.StatusOK, workOrders)
}

func CreateWorkOrder(c *gin.Context) {
	var workOrder models.WorkOrder
	if err := c.ShouldBindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
