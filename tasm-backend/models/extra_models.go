package models

import (
	"gorm.io/gorm"
	"time"
)

type AuditSession struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Title            string         `json:"title"`
	StartDate        time.Time      `json:"startDate"`
	EndDate          *time.Time     `json:"endDate"`
	Status           string         `json:"status"` // e.g. Active, Completed
	TotalAssets      int            `json:"totalAssets"`
	ScannedAssets    int            `json:"scannedAssets"`
	DiscrepancyCount int            `json:"discrepancyCount"`
	AuditorName      string         `json:"auditorName"`
	Progress         int            `json:"progress"`
}

type Consumable struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"name"`
	Category     string         `json:"category"`
	CurrentStock int            `json:"currentStock"`
	ReorderLevel int            `json:"reorderLevel"`
	Location     string         `json:"location"`
}

type MaintenanceContract struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	VendorName  string         `json:"vendorName"`
	ServiceType string         `json:"serviceType"`
	StartDate   time.Time      `json:"startDate"`
	EndDate     time.Time      `json:"endDate"`
	Status      string         `json:"status"` // Active, Expired
}

type WorkOrder struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	WorkOrderID   string         `json:"workOrderId"`
	Title         string         `json:"title"`
	AssetLocation string         `json:"assetLocation"`
	Issue         string         `json:"issue"`
	Severity      string         `json:"severity"` // Critical, High, Medium, Low
	TargetDate    time.Time      `json:"targetDate"`
	Status        string         `json:"status"` // Open, In Progress, Closed
	Technician    string         `json:"technician"`
	Cost          float64        `json:"cost"`
}

type ProcurementRequest struct {
	ID                uint           `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	Title             string         `json:"title"`
	Status            string         `json:"status"`   // Draft, Pending Approval, PO Raised, Shipping, Received
	Priority          string         `json:"priority"` // Low, Medium, High
	EstimatedValue    float64        `json:"estimatedValue"`
	ActualValue       float64        `json:"actualValue"`
	RequestorName     string         `json:"requestorName"`
	RequestorInitials string         `json:"requestorInitials"`
	Department        string         `json:"department"`
	PONumber          string         `json:"poNumber"`
}

type LedgerEntry struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	TransactionID string         `json:"transactionId"`
	Date          time.Time      `json:"date"`
	Description   string         `json:"description"`
	Amount        float64        `json:"amount"`
	Type          string         `json:"type"` // Credit, Debit
	Category      string         `json:"category"`
}

type LeaseAgreement struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	LeaseID     string         `json:"leaseId"`
	Vendor      string         `json:"vendor"`
	StartDate   time.Time      `json:"startDate"`
	EndDate     time.Time      `json:"endDate"`
	MonthlyCost float64        `json:"monthlyCost"`
	Status      string         `json:"status"` // Active, Expired
}

type DepreciationSchedule struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	AssetID       string         `json:"assetId"`
	AssetName     string         `json:"assetName"`
	PurchaseValue float64        `json:"purchaseValue"`
	CurrentValue  float64        `json:"currentValue"`
	Method        string         `json:"method"` // Straight Line, Declining Balance
}

type SoftwareLicense struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	SoftwareName string         `json:"softwareName"`
	PlanName     string         `json:"planName"`
	Status       string         `json:"status"` // Active, Expiring Soon, Expired
	TotalSeats   int            `json:"totalSeats"`
	UsedSeats    int            `json:"usedSeats"`
	RenewalDate  time.Time      `json:"renewalDate"`
	AnnualCost   float64        `json:"annualCost"`
}

// User Management Models
type SystemUser struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	EmployeeID   string         `json:"employeeId"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	PasswordHash string         `json:"-"`
	Department   string         `json:"department"`
	Role         string         `json:"role"`
	Status       string         `json:"status"` // Active, Inactive
	LastLogin    time.Time      `json:"lastLogin"`
}

type AuditDiscrepancy struct {
	ID                uint           `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	AssetTag          string         `json:"assetTag"`
	IssueType         string         `json:"issueType"` // Missing, Location Mismatch, Unregistered
	LastKnownLocation string         `json:"lastKnownLocation"`
	ScannedLocation   string         `json:"scannedLocation"`
	RecommendedAction string         `json:"recommendedAction"`
}

type Reservation struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string         `json:"title"`
	Type      string         `json:"type"`   // Meeting Room, Vehicle, AV Equipment
	Status    string         `json:"status"` // Active, Upcoming
	StartTime time.Time      `json:"startTime"`
	EndTime   time.Time      `json:"endTime"`
	Location  string         `json:"location"`
	BookedBy  string         `json:"bookedBy"`
}

type Location struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
	Type      string         `json:"type"` // Office, Data Center, Warehouse
	Address   string         `json:"address"`
	Capacity  int            `json:"capacity"`
	Status    string         `json:"status"` // Active, Inactive
}

type UserRole struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	RoleName    string         `json:"roleName"`
	Description string         `json:"description"`
	UsersCount  int            `json:"usersCount"`
}

type SystemAlert struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string         `json:"title"`
	Message   string         `json:"message"`
	Type      string         `json:"type"` // Critical, Warning, Info
	Source    string         `json:"source"`
	IsRead    bool           `json:"isRead"`
}
