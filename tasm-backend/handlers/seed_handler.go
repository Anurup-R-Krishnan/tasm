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

	// Wipe all tables
	database.DB.Exec("TRUNCATE TABLE assets, audit_sessions, audit_discrepancies, consumables, maintenance_contracts, work_orders, procurement_requests, ledger_entries, lease_agreements, depreciation_schedules, software_licenses, system_users, user_roles, reservations, locations, system_alerts, system_configs RESTART IDENTITY CASCADE")

	// 1. Create Roles
	roles := []models.UserRole{
		{RoleName: "System Admin", Description: "Full administrative access to all campus resources and security settings.", UsersCount: 1},
		{RoleName: "Finance Manager", Description: "Responsible for asset valuation, procurement approvals, and ledger management.", UsersCount: 2},
		{RoleName: "Facilities Head", Description: "Oversees physical infrastructure, maintenance schedules, and audit sessions.", UsersCount: 1},
		{RoleName: "Operations Tech", Description: "Field staff responsible for work orders, asset tagging, and physical audits.", UsersCount: 3},
		{RoleName: "Employee", Description: "Standard staff access for equipment requests and self-service portal.", UsersCount: 50},
	}
	for _, role := range roles {
		database.DB.Create(&role)
	}

	// 2. Create Locations (Technopark Phase-specific)
	locations := []models.Location{
		{Name: "Thejaswini Hall", Type: "Office", Address: "Phase 1, Technopark", Capacity: 500, Status: "Active"},
		{Name: "Nila Building", Type: "Office", Address: "Phase 1, Technopark", Capacity: 400, Status: "Active"},
		{Name: "Ganga Building", Type: "Office", Address: "Phase 3, Technopark", Capacity: 600, Status: "Active"},
		{Name: "Yamuna Building", Type: "Office", Address: "Phase 3, Technopark", Capacity: 600, Status: "Active"},
		{Name: "Data Center A", Type: "Data Center", Address: "Park Centre Basement", Capacity: 50, Status: "Active"},
		{Name: "Maintenance Shed 1", Type: "Warehouse", Address: "Phase 2 Perimeter", Capacity: 100, Status: "Active"},
	}
	for _, loc := range locations {
		database.DB.Create(&loc)
	}

	// 3. Create System Users
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	users := []models.SystemUser{
		{
			EmployeeID: "ADM-001", Name: "Anurag Sharma", Email: "admin@tasm.com",
			PasswordHash: string(hashedPassword), Role: "System Admin", Status: "Active",
			Department: "IT Administration", LastLogin: time.Now(),
		},
		{
			EmployeeID: "FIN-001", Name: "Priya Nair", Email: "priya.n@tasm.com",
			PasswordHash: string(hashedPassword), Role: "Finance Manager", Status: "Active",
			Department: "Finance & Accounts", LastLogin: time.Now().Add(-24 * time.Hour),
		},
		{
			EmployeeID: "FAC-001", Name: "Rajesh Kumar", Email: "rajesh.k@tasm.com",
			PasswordHash: string(hashedPassword), Role: "Facilities Head", Status: "Active",
			Department: "Operations", LastLogin: time.Now().Add(-48 * time.Hour),
		},
		{
			EmployeeID: "TEC-001", Name: "Meera Pillai", Email: "meera.p@tasm.com",
			PasswordHash: string(hashedPassword), Role: "Operations Tech", Status: "Active",
			Department: "IT Support", LastLogin: time.Now().Add(-2 * time.Hour),
		},
	}
	for _, user := range users {
		database.DB.Create(&user)
	}

	// 4. Create Core Assets
	assets := []models.Asset{
		{
			TagID: "TP-LAP-001", Name: "MacBook Pro M2", Category: "Laptops",
			Status: "Assigned", Location: "Thejaswini Hall", Department: "Engineering",
			Value: 185000, PurchaseDate: time.Now().AddDate(-1, 0, 0),
		},
		{
			TagID: "TP-LAP-002", Name: "Dell Latitude 5430", Category: "Laptops",
			Status: "In Stock", Location: "Nila Building", Department: "Sales",
			Value: 95000, PurchaseDate: time.Now().AddDate(0, -6, 0),
		},
		{
			TagID: "TP-SRV-001", Name: "HP ProLiant Gen10", Category: "Servers",
			Status: "Active", Location: "Data Center A", Department: "IT Infrastructure",
			Value: 450000, PurchaseDate: time.Now().AddDate(-2, 0, 0),
		},
		{
			TagID: "TP-HVAC-001", Name: "BlueStar Industrial Chiller", Category: "Infrastructure",
			Status: "Maintenance", Location: "Ganga Building", Department: "Facilities",
			Value: 1250000, PurchaseDate: time.Now().AddDate(-3, 0, 0),
		},
	}
	for _, asset := range assets {
		database.DB.Create(&asset)
	}

	// 5. Create Work Orders
	workOrders := []models.WorkOrder{
		{
			WorkOrderID: "WO-2026-001", Title: "HVAC Compressor Failure",
			AssetLocation: "Ganga Building - Roof", AssetName: "BlueStar Chiller", AssetTag: "TP-HVAC-001",
			Description: "Unit tripped main breaker. Compressor failing to start. Requires urgent inspection.",
			Issue:       "Mechanical Failure", Severity: "Critical", Status: "In Progress",
			ReportedBy: "Rajesh Kumar", Technician: "Michael Reed",
			TargetDate: time.Now().Add(24 * time.Hour), EstimatedCost: 45000,
		},
		{
			WorkOrderID: "WO-2026-002", Title: "Network Rack Cable Management",
			AssetLocation: "Data Center A", AssetName: "HP ProLiant Rack", AssetTag: "TP-SRV-001",
			Description: "Cleanup of patch cables following recent server migration.",
			Issue:       "Maintenance", Severity: "Low", Status: "Open",
			ReportedBy: "Meera Pillai", Technician: "Suresh Mani",
			TargetDate: time.Now().Add(72 * time.Hour), EstimatedCost: 5000,
		},
	}
	for _, wo := range workOrders {
		database.DB.Create(&wo)
	}

	// 6. Create Lease Agreements
	lease := models.LeaseAgreement{
		LeaseID:     "LS-4055",
		Vendor:      "Tata Elxsi",
		StartDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
		MonthlyCost: 450000,
		Status:      "Active",
	}
	database.DB.Create(&lease)

	// Link lease to an asset (Server)
	database.DB.Model(&models.Asset{}).Where("tag_id = ?", "TP-SRV-001").Update("lease_id", lease.ID)

	// 7. Create Software Licenses
	licenses := []models.SoftwareLicense{
		{SoftwareName: "Adobe Creative Cloud", PlanName: "Enterprise", Status: "Active", TotalSeats: 50, UsedSeats: 42, RenewalDate: time.Now().AddDate(0, 3, 0), AnnualCost: 125000},
		{SoftwareName: "Microsoft 365", PlanName: "E5 Business", Status: "Active", TotalSeats: 500, UsedSeats: 485, RenewalDate: time.Now().AddDate(0, 8, 0), AnnualCost: 850000},
		{SoftwareName: "Slack", PlanName: "Enterprise Grid", Status: "Expiring Soon", TotalSeats: 200, UsedSeats: 198, RenewalDate: time.Now().Add(15 * 24 * time.Hour), AnnualCost: 45000},
	}
	for _, lic := range licenses {
		database.DB.Create(&lic)
	}

	// 7. System Config
	config := models.SystemConfig{
		Key:   "is_setup_completed",
		Value: "true",
	}
	database.DB.Create(&config)

	// Add a few alerts
	alerts := []models.SystemAlert{
		{Title: "Unassigned Assets Detected", Message: "14 laptops in Phase 3 storage have no assigned custodian.", Type: "Warning", Source: "Audit System", IsRead: false},
		{Title: "Maintenance Contract Expiring", Message: "Generator AMC with Kirloskar expires in 12 days.", Type: "Critical", Source: "Finance", IsRead: false},
	}
	for _, alert := range alerts {
		database.DB.Create(&alert)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Database successfully seeded with Technopark Kerala dataset.",
		"admin":   "admin@tasm.com / admin123",
	})
}
