package handlers

import (
	"net/http"
	"tasm-backend/database"
	"tasm-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SeedDatabase(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	// From procurement_handler.go
	requests := []models.ProcurementRequest{
		{Title: "New Employee Laptops (x10)", Status: "Pending Approval", Priority: "High", EstimatedValue: 15000, RequestorName: "Alice Smith", Department: "IT", PONumber: "PO-2023-001"},
		{Title: "Standing Desks (x5)", Status: "PO Raised", Priority: "Medium", EstimatedValue: 2500, RequestorName: "Bob Johnson", Department: "HR", PONumber: "PO-2023-002"},
		{Title: "Office Supplies Q3", Status: "Received", Priority: "Low", EstimatedValue: 500, RequestorName: "Carol Williams", Department: "Admin", PONumber: "PO-2023-003"},
	}
	for _, req := range requests {
		database.DB.Create(&req)
	}
	database.DB.Order("id desc").Find(&requests)
	// From audit_handler.go
	pastDate := time.Now().AddDate(0, -1, 0)
	audits := []models.AuditSession{
		{Title: "Q3 IT Asset Audit", StartDate: time.Now().AddDate(0, 0, -5), Status: "Active", TotalAssets: 450, ScannedAssets: 320, DiscrepancyCount: 12, AuditorName: "Sarah Jenkins", Progress: 71},
		{Title: "Annual Compliance 2024", StartDate: pastDate.AddDate(0, 0, -15), EndDate: &pastDate, Status: "Completed", TotalAssets: 1200, ScannedAssets: 1200, DiscrepancyCount: 5, AuditorName: "Mike Ross", Progress: 100},
		{Title: "Furniture Inventory Audit", StartDate: time.Now().AddDate(0, 0, -1), Status: "Active", TotalAssets: 800, ScannedAssets: 150, DiscrepancyCount: 2, AuditorName: "Emma Clark", Progress: 18},
	}
	for _, a := range audits {
		database.DB.Create(&a)
	}
	database.DB.Order("id desc").Find(&audits)
	// From audit_handler.go
	discs := []models.AuditDiscrepancy{
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
	// From user_handler.go
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	users := []models.SystemUser{
		{
			EmployeeID:   "EMP-001",
			Name:         "Rahul Menon",
			Email:        "rahul.m@technopark.org",
			PasswordHash: string(hashedPassword),
			Department:   "IT Operations",
			Role:         "System Admin",
			Status:       "Active",
			LastLogin:    time.Now(),
		},
		{
			EmployeeID:   "EMP-042",
			Name:         "Priya Suresh",
			Email:        "priya.s@technopark.org",
			PasswordHash: string(hashedPassword),
			Department:   "Finance",
			Role:         "Finance Manager",
			Status:       "Active",
			LastLogin:    time.Now().Add(-2 * time.Hour),
		},
		{
			EmployeeID:   "EMP-088",
			Name:         "Anil Kumar",
			Email:        "anil.k@technopark.org",
			PasswordHash: string(hashedPassword),
			Department:   "Facilities",
			Role:         "Facilities Head",
			Status:       "Inactive",
			LastLogin:    time.Now().Add(-48 * time.Hour),
		},
	}
	for _, u := range users {
		database.DB.Create(&u)
	}
	database.DB.Find(&users)
	// From user_handler.go
	roles := []models.UserRole{
		{
			RoleName:    "System Admin",
			Description: "Full access to all system modules and configuration.",
			UsersCount:  3,
		},
		{
			RoleName:    "Finance Manager",
			Description: "Access to ledgers, depreciation, and lease agreements.",
			UsersCount:  5,
		},
		{
			RoleName:    "Facilities Head",
			Description: "Access to maintenance, work orders, and asset registry.",
			UsersCount:  8,
		},
		{
			RoleName:    "Employee",
			Description: "Self-service portal access for checking out assets.",
			UsersCount:  245,
		},
	}
	for _, r := range roles {
		database.DB.Create(&r)
	}
	database.DB.Find(&roles)
	// From maintenance_handler.go
	contracts := []models.MaintenanceContract{
		{
			VendorName:  "TechServe Solutions",
			ServiceType: "IT Equipment Servicing",
			StartDate:   time.Now().AddDate(0, -6, 0),
			EndDate:     time.Now().AddDate(0, 6, 0),
			Status:      "Active",
		},
		{
			VendorName:  "Cooling Systems Inc.",
			ServiceType: "HVAC Maintenance",
			StartDate:   time.Now().AddDate(-1, 0, 0),
			EndDate:     time.Now().AddDate(0, -1, 0),
			Status:      "Expired",
		},
	}
	for _, contract := range contracts {
		database.DB.Create(&contract)
	}
	database.DB.Order("start_date desc").Find(&contracts)
	// From maintenance_handler.go
	workOrders := []models.WorkOrder{
		{
			WorkOrderID:   "WO-2023-001",
			Title:         "Fix AC in Server Room",
			AssetLocation: "Server Room A",
			Issue:         "AC unit is leaking water",
			Severity:      "High",
			TargetDate:    time.Now().AddDate(0, 0, 2),
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
			TargetDate:    time.Now().AddDate(0, 1, 0),
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
			TargetDate:    time.Now().AddDate(0, 0, -7),
			Status:        "Closed",
			Technician:    "Suresh M.",
			Cost:          3200.00,
		},
	}
	for _, wo := range workOrders {
		database.DB.Create(&wo)
	}
	database.DB.Order("target_date asc").Find(&workOrders)
	// From reservation_handler.go
	nowTime := time.Now()
	reservations := []models.Reservation{
		{Title: "Nila 102", Type: "Meeting Room", Status: "Active", StartTime: nowTime.Add(-time.Hour), EndTime: nowTime.Add(time.Hour), Location: "Nila Building", BookedBy: "Alice Johnson"},
		{Title: "Epson Projector Pro", Type: "AV Equipment", Status: "Upcoming", StartTime: nowTime.Add(2 * time.Hour), EndTime: nowTime.Add(4 * time.Hour), Location: "IT Store", BookedBy: "Bob Smith"},
		{Title: "Corporate Van", Type: "Vehicle", Status: "Upcoming", StartTime: nowTime.AddDate(0, 0, 1), EndTime: nowTime.AddDate(0, 0, 2), Location: "Parking Basement", BookedBy: "Charlie Davis"},
	}
	for _, r := range reservations {
		database.DB.Create(&r)
	}
	database.DB.Order("id desc").Find(&reservations)
	// From alert_handler.go
	alerts := []models.SystemAlert{
		{Title: "Warranty Expired", Message: "Main Generator B (GEN-B-002) warranty has expired. Immediate renewal required.", Type: "Critical", Source: "Asset Management", IsRead: false},
		{Title: "Maintenance Due", Message: "HVAC Unit 4 (HVAC-004) scheduled maintenance is due in 3 days.", Type: "Warning", Source: "Maintenance", IsRead: false},
		{Title: "Stock Low", Message: "Server Rack Screws inventory dropping below threshold (current: 45, min: 50).", Type: "Info", Source: "Inventory", IsRead: false},
	}
	for _, a := range alerts {
		database.DB.Create(&a)
	}
	database.DB.Order("id desc").Find(&alerts)
	// From software_license_handler.go
	licenses := []models.SoftwareLicense{
		{SoftwareName: "Adobe Creative Cloud", PlanName: "Enterprise All Apps", Status: "Active", TotalSeats: 100, UsedSeats: 85, RenewalDate: time.Now().AddDate(0, 6, 0), AnnualCost: 84000},
		{SoftwareName: "Microsoft 365", PlanName: "E5 License", Status: "Active", TotalSeats: 500, UsedSeats: 475, RenewalDate: time.Now().AddDate(0, 10, 0), AnnualCost: 228000},
		{SoftwareName: "Figma", PlanName: "Organization Plan", Status: "Expiring Soon", TotalSeats: 50, UsedSeats: 50, RenewalDate: time.Now().AddDate(0, 0, 15), AnnualCost: 27000},
	}
	for _, l := range licenses {
		database.DB.Create(&l)
	}
	database.DB.Order("id asc").Find(&licenses)
	// From consumable_handler.go
	consumables := []models.Consumable{
		{
			Name:         "Printer Paper A4",
			Category:     "Office Supplies",
			CurrentStock: 45,
			ReorderLevel: 20,
			Location:     "Supply Room 1",
		},
		{
			Name:         "Black Ink Cartridge (HP 64)",
			Category:     "Printer Supplies",
			CurrentStock: 5,
			ReorderLevel: 10,
			Location:     "IT Storage",
		},
		{
			Name:         "Whiteboard Markers (Pack of 4)",
			Category:     "Office Supplies",
			CurrentStock: 30,
			ReorderLevel: 15,
			Location:     "Supply Room 2",
		},
		{
			Name:         "AAA Batteries (Pack of 12)",
			Category:     "Electronics",
			CurrentStock: 12,
			ReorderLevel: 25,
			Location:     "IT Storage",
		},
	}
	for _, cItem := range consumables {
		database.DB.Create(&cItem)
	}
	database.DB.Order("name asc").Find(&consumables)
	// From financial_handler.go
	ledgers := []models.LedgerEntry{
		{TransactionID: "TRX-1001", Date: time.Now().AddDate(0, 0, -2), Description: "Server Rack Purchase", Amount: 15000, Type: "Debit", Category: "Hardware"},
		{TransactionID: "TRX-1002", Date: time.Now().AddDate(0, 0, -5), Description: "Software License Renewal", Amount: 3500, Type: "Debit", Category: "Software"},
		{TransactionID: "TRX-1003", Date: time.Now().AddDate(0, 0, -10), Description: "Old Asset Sale", Amount: 2000, Type: "Credit", Category: "Hardware"},
		{TransactionID: "TRX-1004", Date: time.Now().AddDate(0, -1, 0), Description: "Consulting Fees", Amount: 5000, Type: "Debit", Category: "Services"},
		{TransactionID: "TRX-1005", Date: time.Now().AddDate(0, -1, -15), Description: "Cloud Hosting", Amount: 1200, Type: "Debit", Category: "Infrastructure"},
	}
	for _, l := range ledgers {
		database.DB.Create(&l)
	}
	database.DB.Order("date desc").Find(&ledgers)
	// From financial_handler.go
	leases := []models.LeaseAgreement{
		{LeaseID: "LSE-2023-01", Vendor: "Dell EMC", StartDate: time.Now().AddDate(-1, 0, 0), EndDate: time.Now().AddDate(2, 0, 0), MonthlyCost: 450.00, Status: "Active"},
		{LeaseID: "LSE-2023-02", Vendor: "Cisco Systems", StartDate: time.Now().AddDate(0, -6, 0), EndDate: time.Now().AddDate(1, 6, 0), MonthlyCost: 1200.00, Status: "Active"},
		{LeaseID: "LSE-2021-01", Vendor: "HP Enterprise", StartDate: time.Now().AddDate(-3, 0, 0), EndDate: time.Now().AddDate(-1, 0, 0), MonthlyCost: 800.00, Status: "Expired"},
	}
	for _, l := range leases {
		database.DB.Create(&l)
	}
	database.DB.Order("start_date desc").Find(&leases)
	// From financial_handler.go
	schedules := []models.DepreciationSchedule{
		{AssetID: "AST-2023-001", AssetName: "Dell PowerEdge Server", PurchaseValue: 5000.00, CurrentValue: 3500.00, Method: "Straight Line"},
		{AssetID: "AST-2023-002", AssetName: "Cisco Core Switch", PurchaseValue: 12000.00, CurrentValue: 9000.00, Method: "Declining Balance"},
		{AssetID: "AST-2022-045", AssetName: "Office Chair batch (50)", PurchaseValue: 2500.00, CurrentValue: 1500.00, Method: "Straight Line"},
		{AssetID: "AST-2021-112", AssetName: "Conference Room Display", PurchaseValue: 3000.00, CurrentValue: 500.00, Method: "Straight Line"},
	}
	for _, s := range schedules {
		database.DB.Create(&s)
	}
	database.DB.Order("id desc").Find(&schedules)
	// From asset_handler.go
	assets := []models.Asset{
		{
			Name:           "MacBook Pro 16\"",
			TagID:          "TAG-1082",
			Category:       "IT Equipment",
			Location:       "Bldg 3, Fl 4",
			Status:         "Checked Out",
			Custodian:      "Sarah Jenkins",
			CurrentStock:   1,
			ReorderLevel:   5,
			Value:          2499.99,
			PurchaseDate:   time.Now().AddDate(-1, -6, 0),
			WarrantyStatus: "Active",
		},
		{
			Name:           "Dell Ultrasharp 27\"",
			TagID:          "TAG-2094",
			Category:       "IT Peripherals",
			Location:       "Stockroom A",
			Status:         "In Stock",
			Custodian:      "IT Admin",
			CurrentStock:   12,
			ReorderLevel:   10,
			Value:          399.99,
			PurchaseDate:   time.Now().AddDate(-2, 0, 0),
			WarrantyStatus: "Expired",
		},
		{
			Name:           "Herman Miller Aeron",
			TagID:          "TAG-5032",
			Category:       "Office Furniture",
			Location:       "Bldg 1, Fl 2",
			Status:         "Reserved",
			Custodian:      "New Hire (Pending)",
			CurrentStock:   5,
			ReorderLevel:   2,
			Value:          1200.00,
			PurchaseDate:   time.Now().AddDate(0, -1, 0),
			WarrantyStatus: "Active",
		},
	}
	for _, a := range assets {
		database.DB.Create(&a)
	}
	database.DB.Order("created_at desc").Find(&assets)
	// From location_handler.go
	locations := []models.Location{
		{Name: "Headquarters", Type: "Office", Address: "123 Main St, New York, NY", Capacity: 500, Status: "Active"},
		{Name: "Data Center Alpha", Type: "Data Center", Address: "45 Tech Park, Austin, TX", Capacity: 50, Status: "Active"},
		{Name: "West Coast Hub", Type: "Office", Address: "900 Market St, San Francisco, CA", Capacity: 200, Status: "Active"},
	}
	for _, l := range locations {
		database.DB.Create(&l)
	}
	database.DB.Order("id desc").Find(&locations)

	c.JSON(http.StatusOK, gin.H{"message": "Database seeded successfully"})
}
